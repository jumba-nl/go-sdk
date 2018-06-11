package model

import (
	"math"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)
//proteus:generate
type Realter struct {
	Name  string
	URL   string
	Phone string
}

type AddressLatlon struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json:"lon" bson:"lng"`
}

type AddressWoningwaarde struct {
	ID                 bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Waardekadasterdata int64         `json:"waarde_kd" bson:"waarde_kd,omitempty"`
}

type AddressValue struct {
	ID string `json:"id" bson:"_id,omitempty"`
}

type AddressEnergyLabel2015 struct {
	ID              bson.ObjectId `json:"id" bson:"_id,omitempty"`
	VBOID           int64         `json:"BAGID" bson:"BAGID"`
	LabelVoorlopig  string        `json:"labelVoorlopig" bson:"labelVoorlopig"`
	LabelDefinitief string        `json:"labelDefinitief" bson:"labelDefinitief"`
	LabelModelmatig string        `json:"labelModelmatig" bson:"labelModelmatig"`
	A               float64       `json:"A" bson:"A"`
	Ap              float64       `json:"A+" bson:"A+"`
	App             float64       `json:"A++" bson:"A++"`
	B               float64       `json:"B" bson:"B"`
	C               float64       `json:"C" bson:"C"`
	D               float64       `json:"D" bson:"D"`
	E               float64       `json:"E" bson:"E"`
	F               float64       `json:"F" bson:"F"`
	G               float64       `json:"G" bson:"G"`
	LabelA          string        `json:"LabelA" bson:"LabelA"`
	LabelB          string        `json:"LabelB" bson:"LabelB"`
	LabelC          string        `json:"LabelC" bson:"LabelC"`
	LabelD          string        `json:"LabelD" bson:"LabelD"`
	LabelE          string        `json:"LabelE" bson:"LabelE"`
	LabelF          string        `json:"LabelF" bson:"LabelF"`
	LabelG          string        `json:"LabelG" bson:"LabelG"`
}

type AddressPricingRegion struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Date   string        `json:"date" bson:"date"`
	Region string        `bson:"region"`
	Value  int64         `bson:"value"`
}

type DetailHousehold struct {
	P4                   string
	Totaal               int64   `bson:"totaalhuishoudens"`
	Eenpersoons          int64   `bson:"totaaleenpersoonshuishoudens"`
	Meerpersoons         int64   `bson:"totaalmeerpersoonshuishoudens"`
	MeerpersoonsKinderen int64   `bson:"totaalmeerpersoonshuishoudenskinderen"`
	Gemiddeld            float64 `bson:"gemiddeldegrootte"`
}


// swagger:model Lasten
type Lasten struct {
	ID                                      bson.ObjectId `bson:"_id,omitempty"`
	Postalcode                              string        `bson:"postalcode"`
	City                                    string        `bson:"city,omitempty"`
	County                                  string        `bson:"county,omitempty"`
	Province                                string        `bson:"province,omitempty"`
	OzbRate                                 float64       `bson:"ozb_tarief_woningen"`
	DogLevy                                 float64       `bson:"hondenbelasting"`
	TouristLevy                             float64       `bson:"toeristenbelasting"`
	WasteCollectionLevyOnePersonHousehold   int64         `bson:"reinigingsheffing_eenpersoonshuishouden"`
	WasteCollectionLevyMultiPersonHousehold int64         `bson:"reinigingsheffing_meerpersoonshuishouden"`
	SewageLevyOnePersonHousehold            int64         `bson:"rioolheffing_eenpersoonshuishouden"`
	SewegeLevyMultiPersonHousehold          int64         `bson:"rioolheffing_meerpersoonshuishouden"`
	AverageYearlyCostsOnePersonHousehold    float64       `bson:"woonlasten_eenpersoonshuishouden"`
	AverageYearlyCostsMultiPersonHousehold  float64       `bson:"woonlasten_meerpersoonshuishouden"`
	AverageMonthlyCostsOnePersonHousehold   float64
	AverageMonthlyCostsMultiPersonHousehold float64
}


func GetHousehold(postcode string) (*DetailHousehold, error) {
	item := &DetailHousehold{}

	err := DB().C("huishoudens").Find(bson.M{"p4": postcode}).One(item)

	return item, err
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	_div := math.Copysign(div, val)
	_roundOn := math.Copysign(roundOn, val)
	if _div >= _roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func GetLasten(postalcode string) (*Lasten, error) {
	item := &Lasten{}
	err := DB().C("lasten2016").Find(bson.M{"postalcode": postalcode}).One(item)
	if err == nil {
		item.AverageMonthlyCostsOnePersonHousehold = item.AverageYearlyCostsOnePersonHousehold / 12
		item.AverageMonthlyCostsMultiPersonHousehold = item.AverageYearlyCostsMultiPersonHousehold / 12

		item.AverageMonthlyCostsOnePersonHousehold = Round(item.AverageMonthlyCostsOnePersonHousehold, .5, 2)
		item.AverageMonthlyCostsMultiPersonHousehold = Round(item.AverageMonthlyCostsMultiPersonHousehold, .5, 2)
	}

	return item, err
}

func GetDetailedLasten(postalcode string, price int64) (*Lasten, error) {
	item := &Lasten{}
	err := DB().C("lasten2016").Find(bson.M{"postalcode": postalcode}).One(item)
	if err == nil {
		ozb := (item.OzbRate / 100) * float64(price)
		graphics, err := GetDemographics(bson.M{"postalcode": postalcode})
		if err == nil {
			item.AverageYearlyCostsOnePersonHousehold = float64(item.WasteCollectionLevyOnePersonHousehold) + (graphics.AverageSizeHousehold * float64(item.SewageLevyOnePersonHousehold))
			item.AverageYearlyCostsMultiPersonHousehold = float64(item.WasteCollectionLevyMultiPersonHousehold) + (graphics.AverageSizeHousehold * float64(item.SewegeLevyMultiPersonHousehold))
		}
		item.AverageYearlyCostsOnePersonHousehold = item.AverageYearlyCostsOnePersonHousehold + ozb
		item.AverageYearlyCostsMultiPersonHousehold = item.AverageYearlyCostsMultiPersonHousehold + ozb

		item.AverageMonthlyCostsOnePersonHousehold = item.AverageYearlyCostsOnePersonHousehold / 12
		item.AverageMonthlyCostsMultiPersonHousehold = item.AverageYearlyCostsMultiPersonHousehold / 12

		item.AverageMonthlyCostsOnePersonHousehold = Round(item.AverageMonthlyCostsOnePersonHousehold, .5, 2)
		item.AverageMonthlyCostsMultiPersonHousehold = Round(item.AverageMonthlyCostsMultiPersonHousehold, .5, 2)
		item.AverageYearlyCostsOnePersonHousehold = Round(item.AverageYearlyCostsOnePersonHousehold, .5, 0)
		item.AverageYearlyCostsMultiPersonHousehold = Round(item.AverageYearlyCostsMultiPersonHousehold, .5, 0)
	}

	return item, err
}

func GetWoningwaarde(pht string) (*AddressWoningwaarde, error) {
	item := &AddressWoningwaarde{}

	err := DB().C("woningwaardes").Find(bson.M{"pht": pht}).One(item)

	return item, err
}

func GetEnergyLabel(id string) (*AddressEnergyLabel2015, error) {
	var item AddressEnergyLabel2015
	vboid, err := strconv.ParseInt(strings.TrimLeft(id, "0"), 10, 64)
	if err != nil {
		return &item, err
	}
	err = DB().C("energylabels").Find(bson.M{"BAGID": vboid}).One(&item)
	return &item, err
}

func GetCBSData(province string) ([]AddressPricingRegion, error) {
	items := []AddressPricingRegion{}

	err := DB().C("cbs-price-region").Find(bson.M{"region": province}).All(&items)

	return items, err
}

func GetJumbaWaarde(number, price int64) (value []int64, quality string) {
	quality = "good"
	var step int64 = 0
	var i int64 = 0

	if number <= 100000 {
		step = 10000
		i = 0
	} else if number <= 200000 {
		step = 20000
		i = 100000
	} else if number <= 300000 {
		step = 30000
		i = 200000
	} else if number <= 400000 {
		step = 40000
		i = 300000
	} else if number <= 500000 {
		step = 50000
		i = 400000
	} else {
		step = 100000
		i = 500000
	}

	tmp := number - i
	div := int64(math.Floor(float64(tmp / step)))
	start := i + (step * div)
	end := i + (step * (div + 1))
	value = []int64{start, end}

	if number > 0 {

		pricediff := price - number
		percdiff := float64(pricediff) / float64(number) * float64(100)

		if percdiff < 10 {
			quality = "good"
		} else if percdiff >= 10 && percdiff < 25 {
			quality = "moderate"
		} else if percdiff >= 25 {
			quality = "bad"
		}
	}
	return value, quality
}
