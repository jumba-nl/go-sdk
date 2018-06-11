package model

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Trainstation struct {
	LongName   string `bson:"name"`
	Code       string
	Street     string
	Postalcode string
	Number     string `bson:",omitempty"`
	City       string
	Country    string
	Distance   float64
}

func GetTrainstations(lat, lon float64, amount int) []Trainstation {
	o1 := bson.M{
		"$geoNear": bson.M{
			"near":               []float64{lat, lon},
			"distanceField":      "distance",
			"spherical":          true,
			"distanceMultiplier": 6378000,
			"num":                amount,
		},
	}

	operations := []bson.M{o1}

	pipe := DB().C("trainstations").Pipe(operations)
	results := []Trainstation{}
	err1 := pipe.All(&results)

	if err1 != nil {
		fmt.Printf("ERROR : %s\n", err1.Error())
	}

	return results

}

type Supermarket struct {
	Name     string
	Distance float64
}

func GetSupermarkets(lat, lon float64, amount int) []Supermarket {
	o1 := bson.M{
		"$geoNear": bson.M{
			"near":               []float64{lat, lon},
			"distanceField":      "distance",
			"spherical":          true,
			"distanceMultiplier": 6378000,
			"num":                amount,
		},
	}

	operations := []bson.M{o1}

	pipe := DB().C("supermarkets").Pipe(operations)
	results := []Supermarket{}
	err1 := pipe.All(&results)

	if err1 != nil {
		fmt.Printf("ERROR : %s\n", err1.Error())
	}

	return results

}

type School struct {
	Name     string
	Type     string
	Website  string
	Distance float64
}

func GetSchools(lat, lon float64, amount int) []School {
	o1 := bson.M{
		"$geoNear": bson.M{
			"near":               []float64{lat, lon},
			"distanceField":      "distance",
			"spherical":          true,
			"distanceMultiplier": 6378000,
			"num":                amount,
		},
	}

	operations := []bson.M{o1}

	pipe := DB().C("schools").Pipe(operations)
	results := []School{}
	err1 := pipe.All(&results)

	if err1 != nil {
		fmt.Printf("ERROR : %s\n", err1.Error())
	}

	return results
}

type Demographic struct {
	Postalcode                    string
	City                          string
	County                        string
	Province                      string
	TotalMaleFemale               int64   `bson:"total_male_female"`
	Male                          int64   `bson:"male" `
	Female                        int64   `bson:"female" `
	TotalHouseholds               int64   `bson:"total_households" `
	OnePersonHousehold            int64   `bson:"one_person_household"`
	MorePersonHouseholdWoChildren int64   `bson:"more_person_household_no_children"`
	MorePersonHouseholdWChildren  int64   `bson:"more_person_household_children"`
	AverageSizeHousehold          float64 `bson:"average_household_size" `
}

type AirQuality struct {
	Postalcode      string
	City            string
	County          string
	Province        string
	TotalPopulation int64   `bson:"total_population"`
	NO2             float64 `bson:"NO2"`
	PM10            float64 `bson:"PM10"`
	DeviNO2         float64 `bson:"devi_NO2"`
	DeviPM10        float64 `bson:"devi_PM10"`
}

type MigrationSurplus struct {
	Postalcode string
	City       string
	County     string
	Province   string
	Surplus    float64
}

type WikiDescription struct {
	City     string
	County   string
	Province string
	Summary  string
}

func GetWikiDescription(query bson.M) (description WikiDescription, err error) {
	err = DB().C("cities_descriptions").Find(query).One(&description)
	return
}

func GetMigrationSurplus(query bson.M) (surplus MigrationSurplus, err error) {
	err = DB().C("migration_surplus").Find(query).One(&surplus)
	return
}

func GetAirQuality(query bson.M) (quality AirQuality, err error) {
	err = DB().C("air_quality").Find(query).One(&quality)
	return
}

func GetDemographics(query bson.M) (demo Demographic, err error) {
	err = DB().C("population").Find(query).One(&demo)
	return
}

func GetDemographicsByCity(query bson.M) (demo Demographic, err error) {
	err = DB().C("population").Pipe([]bson.M{bson.M{"$match": query}, bson.M{
		"$group": bson.M{"_id": "$city",
			"city":                              bson.M{"$first": "$city"},
			"total_male_female":                 bson.M{"$sum": "$total_male_female"},
			"male":                              bson.M{"$sum": "$male"},
			"female":                            bson.M{"$sum": "$female"},
			"total_households":                  bson.M{"$sum": "$total_households"},
			"one_person_household":              bson.M{"$sum": "$one_person_household"},
			"more_person_household_no_children": bson.M{"$sum": "$more_person_household_no_children"},
			"more_person_household_children":    bson.M{"$sum": "$more_person_household_children"},
			"average_household_size":            bson.M{"$avg": "$average_household_size"}},
	}}).One(&demo)
	return
}
