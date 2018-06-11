package model

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Offer struct
type Offer struct {
	ID          bson.ObjectId  `bson:"_id,omitempty"`
	Profile     string         `json:"profile,omitempty"`
	FullProfile *ProfileSimple `json:"fullProfile,omitempty"`
	Address     string         `json:"address,omitempty"`
	Offer       int64          `json:"offer,omitempty"`

	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
}

// offers get response struct
// swagger:model OffersGetResponse
type OffersGetResponse struct {
	Total     int
	Limit     int    `json:",omitempty"`
	Offset    int    `json:",omitempty"`
	Sort      string `json:",omitempty"`
	Direction string `json:",omitempty"`
	Filter    bson.M
	Items     []*Offer
}

// offers post response struct
// swagger:model OffersPostResponse
type OffersPostResponse struct {
	Item *Offer
}

// swagger:parameters offersGet
type OffersGetRequest struct {
	// What is the maximum amount of offers that should be returned.
	Limit int `json:"limit"`
	// What is offset from zero to return.
	Offset int `json:"offset"`
	// What value should be used to sort.
	Sort string `json:"sort"`
	// In what direction should be sorted.
	Direction string `json:"direction"`
}

// swagger:parameters offersPost
type OffersPostBody struct {
	// in: body
	Body OffersPostRequest `json:"body"`
}

// swagger:model
type OffersPostRequest struct {
	// Postalcode + number (+ additions).
	// required: true
	Address string `json:"address"`
	// Amount to offer.
	// required: true
	Offer int64 `json:"offer"`
}

// create a single offer
func CreateOffer(o *Offer) (*Offer, error) {
	if !o.ID.Valid() {
		o.ID = bson.NewObjectId()
	}

	_, err := DB().C("jumba_offers").UpsertId(o.ID, o)
	offer, _ := GetOffer(o.ID)

	return offer, err
}

// get a single offer related to an id
func GetOffer(i bson.ObjectId) (*Offer, error) {
	offer := &Offer{}

	err := DB().C("jumba_offers").FindId(i).One(offer)

	return offer, err
}

// get offers by address
func GetOfferByProperty(property string) (*Offer, error) {
	offer := &Offer{}

	err := DB().C("jumba_offers").Find(bson.M{"address": property}).One(offer)

	return offer, err
}

func IsValidOffer(offer int64, address Search) error {
	var err error

	min := 0.8
	max := 1.2
	bid := float64(offer)

	if address.Payload.Price != 0 {
		if bid < (float64(address.Payload.Price) * min) {
			err = errors.New("Offer may only be a maximum of 20% lower than the asking price")
		} else if bid > (float64(address.Payload.Price) * max) {
			err = errors.New("Offer may only be a maximum of 20% higher than the asking price")
		}
	} else if len(address.Payload.JumbaValue) >= 1 && address.Payload.JumbaValue[0] != 0 {
		if bid < (float64(address.Payload.JumbaValue[0]) * min) {
			err = errors.New("Offer may only be a maximum of 20% lower than the lowest price in the Jumba Value range")
		} else if bid > (float64(address.Payload.JumbaValue[1]) * max) {
			err = errors.New("Offer may only be a maximum of 20% higher than the highest price in the Jumba Value range")
		}
	}

	if err != nil {
		return err
	} else {
		return nil
	}
}

func (c *Offer) Save() error {
	if !c.ID.Valid() {
		c.ID = bson.NewObjectId()
	}

	_, err := DB().C("jumba_offers").UpsertId(c.ID, c)
	return err
}

func (c *Offer) Delete() error {
	err := DB().C("jumba_offers").RemoveId(c.ID)
	return err
}
