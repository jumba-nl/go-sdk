package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type News struct {
	ID             bson.ObjectId `bson:"_id"`
	Active         bool
	Slug           string
	SeoTitle       string `bson:"seo-title"`
	SeoDescription string `bson:"seo-description"`
	Image          string
	Title          string
	Description    string
	Author         NewsAuthor `bson:",omitempty" json:",omitempty"`
	Created        time.Time
	Updated        time.Time
}

type NewsAuthor struct {
	ID    string `bson:",omitempty" json:",omitempty"`
	Name  string `bson:",omitempty" json:",omitempty"`
	Image string `bson:",omitempty" json:",omitempty"`
}

type NewsHeader struct {
	Slug    string
	Title   string
	Image   string
	Created time.Time
}

func GetLatestNewsHeaders() (headers []NewsHeader) {
	DB().C("news").Find(bson.M{"active": true}).Sort("-created").Select(bson.M{"title": 1, "slug": 1, "image": 1, "created": 1}).Limit(5).All(&headers)
	return
}

func GetNews(active bool) (news []News) {
	find := bson.M{}
	if active {
		find = bson.M{"active": true}
	}

	news = make([]News, 0)
	DB().C("news").Find(find).Sort("-created").All(&news)
	return
}

func GetNewsArticle(query bson.M) (news News, err error) {
	err = DB().C("news").Find(query).One(&news)
	return
}
