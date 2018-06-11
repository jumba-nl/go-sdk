package model

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// swagger:model Favorite
type Favorite struct {
	ID       bson.ObjectId `bson:"_id"`
	Profile  string        `bson:"profile"`
	Url      string        `json:",omitempty" bson:"url,omitempty"`
	Property string        `bson:"property"`
	Active   bool          `bson:"active"`
	Created  time.Time     `bson:"created"`
	Updated  time.Time     `bson:"updated"`
}

type OldApiFavorite struct {
	ID      bson.ObjectId `bson:"_id"`
	Profile bson.ObjectId `bson:"profile"`
	Url     string        `bson:"url"`
	Created struct {
		Sec  int32
		Usec int32
	} `bson:"created"`
	Updated struct {
		Sec  int32
		Usec int32
	} `bson:"updated"`
}

func GetIdsFromFavorites(favs []Favorite) (ids []string) {
	for _, val := range favs {
		ids = append(ids, val.Property)
	}

	return ids
}

func FetchFavorite(property, profile string) (fav Favorite, err error) {
	err = DB().C("favorites").Find(bson.M{"property": property, "profile": profile}).One(&fav)
	if err != nil && bson.IsObjectIdHex(profile) {
		// Convert old favorites to new favorites for backwards compatibility
		old := OldApiFavorite{}
		query := bson.M{"url": fmt.Sprintf("%s%s", "http://jumba.nl/", property), "profile": bson.ObjectIdHex(profile)}
		if err = DB().C("favorites").Find(query).One(&old); err != nil {
			return fav, err
		} else {
			// remove old favorite
			DB().C("favorites").Remove(query)
		}

		if t, err := time.Parse(time.RFC3339, time.Unix(int64(old.Updated.Sec), 0).Format(time.RFC3339)); err == nil {
			fav.Updated = t
		} else {
			fav.Updated = time.Now()
		}

		if t, err := time.Parse(time.RFC3339, time.Unix(int64(old.Created.Sec), 0).Format(time.RFC3339)); err == nil {
			fav.Created = t
		} else {
			fav.Created = time.Now()
		}

		fav.Active = true
		fav.Property = strings.Replace(old.Url, "http://jumba.nl/", "", 1)
		fav.Profile = profile

		fav.Save()
	}

	return fav, err
}

func FetchFavorites(profile string, all bool) (favs []Favorite, err error) {
	selectors := []bson.M{{"profile": profile}}
	if bson.IsObjectIdHex(profile) {
		selectors = append(selectors, bson.M{"profile": bson.ObjectIdHex(profile)})
	}
	if !all {
		for i := range selectors {
			selectors[i]["active"] = true
		}
	}
	err = DB().C("favorites").Find(bson.M{"$or": selectors}).All(&favs)

	// clean up old values in parallel
	duplicates := false
	sem := make(chan Favorite, len(favs))
	for index, val := range favs {
		if val.Url != "" {
			duplicates = true
			go func(index int, val Favorite) {
				DB().C("favorites").Remove(bson.M{"profile": profile, "url": val.Url})

				favs[index].Active = true
				favs[index].Profile = profile
				favs[index].Property = strings.Replace(val.Url, "http://jumba.nl/", "", 1)
				favs[index].Url = ""
				favs[index].Updated = time.Now()
				favs[index].Created = time.Now()

				exists, _ := DB().C("favorites").Find(bson.M{"profile": profile}).Count()
				if exists == 0 {
					favs[index].Save()
				}

				sem <- val
			}(index, val)
		} else {
			sem <- val
		}
	}

	for i := 0; i < len(favs); i++ {
		<-sem
	}

	if duplicates {
		old := favs
		favs = []Favorite{}
		for _, val := range old {
			if !ExistsInFavoritesSlice(val, favs) {
				favs = append(favs, val)
			}
		}
	}

	return favs, err
}

func ExistsInFavoritesSlice(f Favorite, s []Favorite) bool {
	for _, b := range s {
		if b.Property == f.Property {
			return true
		}
	}
	return false
}

func (f *Favorite) Save() error {
	if !f.ID.Valid() {
		f.ID = bson.NewObjectId()
	}

	_, err := DB().C("favorites").UpsertId(f.ID, bson.M{"$set": f})

	return err
}
