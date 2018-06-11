package model

import (
	"log"
	"strings"
	"time"

	. "github.com/jumba-nl/jumba/ext/config"
	cache "github.com/patrickmn/go-cache"

	"gopkg.in/mgo.v2/bson"
)

const detailCountersCol = "detail_counters"

type VisitCounter []time.Time
type VisitCounterResponse struct {
	ID       bson.ObjectId `bson:"_id" json:",omitempty"`
	Property string
	Total    int64
	Counter  VisitCounter
}

type VisitCounterIdent struct {
	ID       bson.ObjectId `bson:"_id" json:",omitempty"`
	Property string
	Counter  VisitCounter
}

func UpdateCounter(id string) {
	var valid bool
	var count VisitCounter

	// rertieve old cached values
	if cached, exists := Config.Caches.Counter.Cache.Get("counter-" + id); exists {
		count, valid = cached.(VisitCounter)
		if !valid {
			log.Print("invalid type assertion in counter\n")
			return
		}
	}

	// update values
	count = append(count, time.Now())
	Config.Caches.Counter.Cache.Set("counter-"+id, count, cache.DefaultExpiration)
}

func OnCounterExpire(key string, value interface{}) {
	count, valid := value.(VisitCounter)
	if valid {
		var prev VisitCounterIdent
		cur := VisitCounterIdent{Property: strings.Replace(key, "counter-", "", 1), Counter: count}
		err := DB().C(detailCountersCol).Find(bson.M{"property": cur.Property}).One(&prev)
		if err == nil {
			cur.ID = prev.ID
			cur.Counter = append(prev.Counter, cur.Counter...)
		} else {
			cur.ID = bson.NewObjectId()
		}

		DB().C(detailCountersCol).UpsertId(cur.ID, cur)
	} else {
		log.Print("invalid type assertion in counter\n")
	}
}

func RetrieveCounter(id string, start time.Time, end time.Time) VisitCounterResponse {
	var total struct {
		ID      bson.ObjectId `bson:"_id"`
		Total   int64
		Counter VisitCounter
	}
	counter := VisitCounterResponse{Property: id}

	err := DB().C(detailCountersCol).Pipe([]bson.M{
		bson.M{
			"$match": bson.M{"property": id},
		},
		bson.M{
			"$unwind": "$counter",
		},
		bson.M{
			"$group": bson.M{"_id": "$_id", "total": bson.M{"$sum": 1}},
		},
	}).One(&total)
	if err == nil {
		counter.Total = total.Total

		err = DB().C(detailCountersCol).Pipe([]bson.M{
			bson.M{
				"$match": bson.M{"property": id},
			},
			bson.M{
				"$unwind": "$counter",
			},
			bson.M{
				"$match": bson.M{"counter": bson.M{"$gte": start}},
			},
			bson.M{
				"$match": bson.M{"counter": bson.M{"$lte": end}},
			},
			bson.M{
				"$group": bson.M{"_id": "$_id", "counter": bson.M{"$push": "$counter"}},
			},
		}).One(&total)
		if err == nil {
			counter.Counter = total.Counter
		} else {
			counter.Counter = make([]time.Time, 0)
		}
	} else {
		counter.Total = 0
		counter.Counter = make([]time.Time, 0)
	}

	return counter
}

func RetrieveWeeklyCounter(id string) int64 {
	cur := VisitCounterIdent{}
	DB().C(detailCountersCol).Pipe([]bson.M{
		bson.M{
			"$match": bson.M{"property": id},
		},
		bson.M{
			"$unwind": "$counter",
		},
		bson.M{
			"$match": bson.M{"counter": bson.M{"$gte": time.Now().AddDate(0, 0, -7)}},
		},
		bson.M{
			"$group": bson.M{"_id": "$_id", "counter": bson.M{"$push": "$counter"}},
		},
	}).One(&cur)

	if cached, exists := Config.Caches.Counter.Cache.Get("counter-" + id); exists {
		count, valid := cached.(VisitCounter)
		if valid {
			cur.Counter = append(cur.Counter, count...)
		}
	}

	return int64(len(cur.Counter))
}
