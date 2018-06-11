package model

import (
	"regexp"
	"strings"
	"unicode"

	"gopkg.in/mgo.v2/bson"
)

// FormatIdentifier attempts to convert given id to a valid property ID.
// The formatted identifier is returned accompanied by a boolean, indicating
// whether given ID is considered valid or not. A valid ID is basically any
// postal code (PC6) followed by something.
func FormatIdentifier(id string) (string, bool) {
	id = TrimAllSpaces(id)

	if len(id) > 6 {
		id = strings.ToUpper(id[0:6]) + id[6:]
	}
	return id, isPostalCode.FindString(id) != ""
}

// Regexp to match inputs prefixed by a valid postal code (PC6).
// Letters in postal codes are, by convention, written in upper case. At this stage the input should
// already passed (or called from) FormatIdentifier(), therefore match is made case sensitive.
var isPostalCode = regexp.MustCompile(`^\d{4}\s?[A-Z]{2}`)

func GetSortString(sort string, direction string) string {
	var directionsort string
	if sort != "" && direction != "" {
		if direction == "desc" {
			directionsort = "-"
		} else if direction == "asc" {
			directionsort = "+"
		}
	} else if sort != "" && direction == "" {
		directionsort = "-"
	}
	directionsort += sort
	return directionsort
}

// also returns distinct count of users that posted
func GetAllForParam(t interface{}, collection string, q bson.M, limit int, offset int, sort string) (err error, count int) {
	if sort != "" {
		err = DB().C(collection).Find(q).Sort(sort).Limit(limit).Skip(offset).All(t)
	} else {
		err = DB().C(collection).Find(q).Limit(limit).Skip(offset).All(t)
	}

	count, _ = DB().C(collection).Find(q).Count()

	return
}

func TrimAllSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
