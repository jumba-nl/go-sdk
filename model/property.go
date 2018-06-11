package model

import (
	"encoding/json"
	"errors"
	"time"

	"log"

	"github.com/jumba-nl/jumba/ext/config"

	"golang.org/x/net/context"

	"gopkg.in/mgo.v2/bson"
	"gopkg.in/olivere/elastic.v5"
)

// swagger:model Status
type Status struct {
	ID                   string
	BAGID                string
	Built                int64
	ForSale              bool
	ForSaleSince         time.Time
	Sold                 bool
	SoldSince            time.Time
	SoldUnderReservation bool
	LastModified         time.Time
	OpenHouseDays        []OpenHouseDate
	Prices               []Price
}

type OpenHouseDay struct {
	ID       bson.ObjectId `bson:"_id"`
	Property string
	Date     OpenHouseDate
}

type OpenHouseDate struct {
	From time.Time
	Till time.Time
}

type PriceHistory struct {
	ID          string
	Prices      []Price
	LastUpdated time.Time
}

type Price struct {
	Date  time.Time `bson:"date"`
	Price int64     `bson:"price"`
}

type SkynetProperty struct {
	ID      string    `bson:"id"`
	Address Search    `bson:"address"`
	Time    time.Time `bson:"time"`
	User    bool      `bson:"user"`
}

// Creates an property model from an Elastic return value
func CreateStatusModel(search Search) Status {
	status := Status{}

	status.ID = search.Payload.ID
	status.BAGID = string(search.Payload.BagVBOID)
	status.ForSale = search.Parameters.Forsale
	status.ForSaleSince = search.ForSaleSince
	status.Sold = search.Parameters.Sold
	status.SoldSince = search.SoldSince

	// status scraped from Funda, only way for now how we can know
	if search.Payload.Properties.SaleStatus == "Verkocht onder voorbehoud" {
		status.SoldUnderReservation = true
	} else {
		status.SoldUnderReservation = false
	}

	status.LastModified = search.LastCheck
	status.Built = search.Payload.Built

	priceHistory, err := GetPriceHistory(status.ID)
	if err != nil {
		log.Println(err)

		// create empty slice for json response
		status.Prices = make([]Price, 0)
	} else {
		status.Prices = priceHistory.Prices
	}

	if status.ForSale && len(status.Prices) >= 1 && status.Prices[len(status.Prices)-1].Price != search.Parameters.Price {
		status.Prices = append(status.Prices, Price{Date: status.ForSaleSince, Price: search.Parameters.Price})
	}

	openHouseDays, err := GetEarliestOpenHouseDay(status.ID)
	if err != nil {
		// just log the error but continue
		log.Println(err)
	}

	if len(openHouseDays) >= 1 {
		for _, val := range openHouseDays {
			status.OpenHouseDays = append(status.OpenHouseDays, val.Date)
		}
	} else {
		// create empty slice for json response
		status.OpenHouseDays = make([]OpenHouseDate, 0)
	}

	return status
}

func FetchProperty(id string) (Search, int64) {
	// trim all spaces to reduce failure rate
	id, _ = FormatIdentifier(id)

	query := elastic.NewBoolQuery().MinimumShouldMatch("1")
	// because ID and BagVBOID is analyzed we can't use a term query
	// this means we'll have to filter the results afterwards
	query = query.Should(
		elastic.NewMatchQuery("Payload.ID", id),
		elastic.NewMatchQuery("Payload.BagVBOID", id),
	)
	response, err := config.Config.ElasticSearch.Client.Search().
		Index(config.Config.ElasticSearch.Index).
		Type(config.Config.ElasticSearch.Type).
		Query(query).
		Do(context.TODO())
	if err != nil {
		log.Println(err)
		return Search{}, 0
	}

	var hits int64
	property := Search{}
	for _, val := range response.Hits.Hits {
		err := json.Unmarshal(*val.Source, &property)
		if err != nil {
			log.Print(err)
			continue
		}

		if property.Payload.ID == id || property.Payload.BagVBOID == maybe(id) {
			hits = int64(1)
			break
		}
	}

	return property, hits
}

func FetchPropertyFromLatlng(lat float64, lng float64) (Search, error) {
	query := elastic.NewGeoDistanceQuery("Payload.Location").
		Lat(lat).
		Lon(lng).
		Distance("5m")

	response, _ := config.Config.ElasticSearch.Client.Search().
		Index(config.Config.ElasticSearch.Index).
		Type(config.Config.ElasticSearch.Type).
		Query(query).
		Do(context.TODO())

	if response != nil && response.Hits.TotalHits > 0 {
		for _, val := range response.Hits.Hits {
			prop := Search{}
			err := json.Unmarshal(*val.Source, &prop)
			if err != nil {
				log.Print(err)
			} else {
				return prop, nil
			}
		}
	}

	return Search{}, errors.New("not found")
}

func GetEarliestOpenHouseDay(id string) ([]OpenHouseDay, error) {
	days := []OpenHouseDay{}
	err := DB().C("open_house_day").Find(bson.M{"property": id, "date.from": bson.M{"$gte": time.Now()}}).Sort("date.from").All(&days)

	return days, err
}

func GetPriceHistory(id string) (PriceHistory, error) {
	history := PriceHistory{}
	err := DB().C("price_history").Pipe(
		[]bson.M{
			{"$match": bson.M{"_id": id}},
			{"$unwind": "$prices"},
			{"$sort": bson.M{"prices.date": 1}},
			{
				"$group": bson.M{
					"_id": "$_id",
					"prices": bson.M{
						"$push": bson.M{
							"date":  "$prices.date",
							"price": "$prices.price",
						},
					},
				},
			},
		}).One(&history)

	return history, err
}

func InsertIntoSkynet(identifier string, doc Search, isUser bool) (err error) {
	s := conn.Copy()
	defer s.Close()

	return s.DB("skynet").C("address_history").Insert(SkynetProperty{
		ID:      identifier,
		Address: doc,
		Time:    time.Now(),
		User:    isUser,
	})
}

// SkynetBulkProcessor creates a channel and starts a go routine which reads Search objects from the channel
// and inserts the objects in bulk. The created channel can has a buffer size is determined by buf parameter where
// size sets max number of operations per bulk. Stop the processor by closing the returned channel.
func SkynetBulkProcessor(buf, size int) (chan<- Search, <-chan struct{}) {
	ch := make(chan Search, buf)
	done := make(chan struct{})
	go func() {
		s := conn.Copy()
		defer s.Close()

		bulk := s.DB("skynet").C("address_history").Bulk()
		queued := 0

		for obj := range ch {
			if queued >= size {
				if _, err := bulk.Run(); err != nil {
					log.Println(err)
				}
				bulk = s.DB("skynet").C("address_history").Bulk()
				queued = 0
			}
			bulk.Insert(SkynetProperty{
				ID:      obj.Payload.ID,
				Address: obj,
				Time:    time.Now(),
				User:    false,
			})
			queued++
		}
		if queued > 0 {
			if _, err := bulk.Run(); err != nil {
				log.Println(err)
			}
		}
		done <- struct{}{}
	}()
	return ch, done
}
