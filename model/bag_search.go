package model

import (
	"encoding/json"
	"time"

	"log"
)

// swagger:model Search
type Search struct {
	Input        SearchInput
	Filter       SearchFilter
	Parameters   SearchParameters
	Payload      Combined
	Output       string
	Type         int
	Private      bool
	NoBag        bool
	ForSaleSince time.Time
	SoldSince    time.Time
	LastCheck    time.Time
}

type searchEncodingHelper Search

// TODO Pick the low hanging fruits, write a test voor deze.
func (obj Search) MarshalJSON() ([]byte, error) {
	helper := searchEncodingHelper(obj)
	if helper.Payload.Outline == nil {
		helper.Payload.Outline = []AddressLatlon{}
	}
	if helper.Payload.JumbaValue == nil {
		helper.Payload.JumbaValue = []int64{}
	}
	if helper.Payload.Images == nil {
		helper.Payload.Images = []string{}
	}
	if helper.Payload.Realter == nil {
		helper.Payload.Realter = &Realter{}
	}
	helper.ForSaleSince = obj.ForSaleSince.Truncate(time.Second)
	helper.SoldSince = obj.SoldSince.Truncate(time.Second)
	helper.LastCheck = obj.LastCheck.Truncate(time.Second)
	return json.Marshal(helper)
}

func (obj Search) ObjectID() string {
	return obj.Payload.ID
}

func (obj Search) Extract(s *Search) error {
	*s = obj // simply copy the current value into the parameters location.
	return nil
}

type SearchFilter struct {
	Street     string
	City       string
	Number     string
	Postcode   string
	Province   string
	County     string
	NumberOnly int64
}

type SearchSmall struct {
	ID           string
	Input        SearchInput
	Parameters   SearchParameters
	Payload      CombinedSmall
	Output       string
	Type         int
	NoBag        bool
	ForSaleSince time.Time
	SoldSince    time.Time
	LastCheck    time.Time
}

type SearchSuggest struct {
	Payload CombinedSuggest
	Output  string
	Type    int
}

type SearchInput struct {
	Fulltext string
	Suggest  []string
	Label    string `bson:"label"`
}

type SearchParameters struct {
	Forsale           bool
	Sold              bool
	AmountForsale     int64
	Price             int64
	Size              int64
	ClusterMembership int64
}

type Combined struct {
	ID                string     `bson:"_id,omitempty"`
	Address           string
	Country           string
	Province          string
	County            string
	Main              *Address
	Related           []*Address `json:",omitempty"`
	Built             int64
	Size              int64
	Function          string
	Forsale           bool
	Sold              bool
	Location          AddressLatlon
	Outline           []AddressLatlon
	Price             int64
	JumbaValue        []int64    `bson:"jumba_value"`
	JumbaValueQuality string     `bson:"jumba_value_quality"`
	Image             string
	Images            []string   `bson:"images"`
	LivingScore       CombinedLivingScores
	Funda             CombinedFunda
	Jaap              CombinedJaap

	Kenmerken mapStringString `bson:"kenmerken"`

	Type                           string `bson:"type"`
	Bedrooms                       int64  `bson:"bedrooms"`
	PerceelSize                    int64  `bson:"percopp"`
	Rooms                          int64  `bson:"rooms"`
	Details                        string `bson:"details"`
	Maintenance                    string `bson:"maintenance"`
	Sanitation                     string `bson:"sanitation"`
	Kitchen                        string `bson:"kitchen"`
	Paintwork                      string `bson:"paintwork"`
	Garden                         string `bson:"garden"`
	View                           string `bson:"view"`
	Balkony                        string `bson:"balkony"`
	Garage                         string `bson:"garage"`
	Isolation                      string `bson:"isolation"`
	Heating                        string `bson:"heating"`
	EnergylabelEstimate            string `bson:"energylabel_estimate"`
	EnergylabelConsumptionEstimate string `bson:"energylabelconsumption_estimate"`
	BagVBOID                       maybe  `bson:"bag_vbo_id"`
	Purpose                        string `bson:"purpose"`

	Realter *Realter

	Properties Properties               `bson:"properties"`
	Floors     []map[string]interface{} `bson:"floors" json:",omitempty"`
	Labels     []string                 `bson:"labels" json:",omitempty"`
	Financial  CombinedFinancial        `bson:"financial" json:",omitempty"`
}

// maybe implements an optional string by ignoring Unmarshal errors.
type maybe string

// UnmarshalJSON attempts to unmarshal given bytes into a string but doesn't raise an error on failure.
// If it fails it just logs the encountered error, leaves the value unchanged and continues as if nothing happened.
func (m *maybe) UnmarshalJSON(b []byte) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		// Just log the error, nothing we can do about this.
		log.Println(err)
	} else {
		*m = maybe(value)
	}
	return nil
}

// String implements the fmt.Stringer interface in an effort to make the maybe type behave
// as much the same as the native string type as possible.
func (m maybe) String() string {
	return string(m)
}

// mapStringString ignores inputs json.Unmarshal cannot handle.
type mapStringString map[string]string

// UnmarshalJSON "swallows" any potential errors returned by json.Unmarshal.
func (mss *mapStringString) UnmarshalJSON(b []byte) error {
	var value map[string]string
	if err := json.Unmarshal(b, &value); err != nil {
		log.Println(err)
	} else {
		*mss = mapStringString(value)
	}
	return nil
}

// Properties holds al form filled values
type Properties struct {
	Bidding                string
	Bouwvorm               string
	AppartementSubType     string
	AppartementType        string
	WoonhuisType           string
	WoonhuisSubType        string
	SaleStatus             string
	Aanvaarding            string
	AanvaardingsDatum      time.Time
	AanvaardingsToevoeging string
	ServiceCosts           float64

	Dimensions struct {
		Inhoud                     int64
		ExterneBergRuimte          int64
		OverigeInpandigeRuimte     int64
		GebouwGebondenBuitenRuimte int64
	}

	Onderhoud struct {
		Binnen string
		Buiten string
	}

	Bijzonderheden struct {
		BeschermdStadsDorpsGezicht bool
		DubbeleBewoningAanwezig    bool
		DubbeleBewoningMogelijk    bool
		GedeeltelijkGestoffeerd    bool
		GedeeltelijkVerhuurd       bool
		Gemeubileerd               bool
		Gestoffeerd                bool
		Monument                   bool
		MonumentaalPand            bool
		ToegankelijkGehandicapten  bool
		ToegankelijkOuderen        bool
	}

	Location struct {
		ForestEdge  bool
		BusyRoad    bool
		Park        bool
		QuietRoad   bool
		FairWay     bool
		Water       bool
		Sheltered   bool
		CountrySide bool
		Forest      bool
		Centre      bool
		Suburb      bool
		OpenSpace   bool
		FreeView    bool
	}

	Garden struct {
		Back        bool
		No          bool
		Patio       bool
		Place       bool
		Around      bool
		Front       bool
		Side        bool
		SunTerrace  bool
		RoofTerrace bool
		Balcony     bool
		Main        string
		MainLength  int64
		MainWidth   int64
		MainSize    int64
		AllSize     int64
		Condition   string
		AroundBack  bool
		Orientation string
	}

	Lavatories int64
	Bathroom struct {
		Amount         int64
		Bathtub        bool
		SittingBathtub bool
		Shower         bool
		Lavatory       bool
		Sauna          bool
	}

	Storage struct {
		Type   string
		Amount string

		Isolation struct {
			Dakisolatie            bool
			DubbelGlas             bool
			EcoBouw                bool
			GedeeltelijkDubbelGlas bool
			GeenIsolatie           bool
			GeenSpouw              bool
			MuurIsolatie           bool
			VloerIsolatie          bool
			VolledigGeisoleerd     bool
			VoorzetRamen           bool
		}

		Options struct {
			ElectricDoor bool
			Electricity  bool
			Loft         bool
			Heating      bool
			Water        bool
		}
	}

	Garage struct {
		AangebouwdHout   bool
		AangebouwdSteen  bool
		Carport          bool
		GarageMetCarport bool
		Box              bool
		Inpandig         bool
		Losstaand        bool
		Aangebouwd       bool
		Souterrain       bool
		VrijstaandHout   bool
		VrijstaandSteen  bool

		Parkeerkelder     bool
		Parkeerplaats     bool
		BetaaldParkeren   bool
		Vergunninghouders bool
		Parkeergarages    bool
		VrijParkeren      bool
		GeenGarage        bool
		GarageMogelijk    bool

		Options struct {
			ElectricDoor bool
			Electricity  bool
			Loft         bool
			Heating      bool
			Water        bool
		}

		Amount string
		Cars   string
		Size   int64
	}

	Description struct {
		Algemeen       string
		Buurt          string
		Bereikbaarheid string
		Faciliteiten   string
		Bijzonderheden string
		LinkMeerInfo   string
	}

	Roof struct {
		Type string
		Material struct {
			Asbest       bool
			Bitumineuze  bool
			Dakpannen    bool
			Kunststof    bool
			Leisteen     bool
			MetaalOverig bool
			Riet         bool
		}
	}

	Isolation struct {
		Dakisolatie            bool
		DubbelGlas             bool
		EcoBouw                bool
		GedeeltelijkDubbelGlas bool
		GeenIsolatie           bool
		GeenSpouw              bool
		MuurIsolatie           bool
		VloerIsolatie          bool
		VolledigGeisoleerd     bool
		VoorzetRamen           bool
	}

	WarmWater struct {
		CentraleVoorziening       bool
		CVKetel                   bool
		ElektrischeBoilerEigendom bool
		ElektrischeBoilerHuur     bool
		GasBoilerEigendom         bool
		GasBoilerHuur             bool
		GeenVerwarmdWater         bool
		GeiserEigendom            bool
		GeiserHuur                bool
		ZonneBoiler               bool
		ZonneCollectoren          bool
	}

	Heating struct {
		Main                        string
		Blokverwarming              bool
		CVKetel                     bool
		ElektrischeVerwarming       bool
		GasKachels                  bool
		GeenVerwarming              bool
		HeteLuchtVerwarming         bool
		KolenKachel                 bool
		MoederHaard                 bool
		MogelijkheidOpenHaard       bool
		MuurVerwarming              bool
		OpenHaard                   bool
		StadsVerwarming             bool
		VloerverwarmingGedeeltelijk bool
		VloerverwarmingGeheel       bool
		Warmtepomp                  bool
		Pelletkachel                bool
	}

	CV struct {
		Type  string
		Built string
		Owned string
		Fuel  string
		Combi string
	}

	Options struct {
		Airconditioning       bool
		Alarm                 bool
		Dakraam               bool
		FransBalkon           bool
		Jacuzzi               bool
		Lift                  bool
		MechanischeVentilatie bool
		Rolluiken             bool
		RookKanaal            bool
		Satellietschotel      bool
		Sauna                 bool
		Schuifpui             bool
		StoomCabine           bool
		TVKabel               bool
		Windmolen             bool
		ZonneCollectoren      bool
		Zonwering             bool
		Zwembad               bool
	}
}

type CombinedSmall struct {
	ID       string `bson:"_id,omitempty"`
	Address  string
	Country  string
	Province string
	County   string
	Main     *Address
}

type CombinedSuggest struct {
	Input    SearchInput
	Image    string
	Province string
	Label    string
	Main     *Address
}

type CombinedFunda struct {
	URL         string
	Price       int64
	Rooms       int64
	Size        int64
	Description string
	Since       string
}

type CombinedJaap struct {
	URL         string
	Price       int64
	Rooms       int64
	Size        int64
	Description string
	Since       string
}

type CombinedLivingScores struct {
	PC4        string  `bson:"pc4"`
	Score      int64   `bson:"score"`
	Year       int64   `bson:"year"`
	DiffPub    float64 `bson:"diff_pub"`
	DiffWon    float64 `bson:"diff_won"`
	DiffVrz    float64 `bson:"diff_vrz"`
	DiffBev    float64 `bson:"diff_bev"`
	DiffLftsam float64 `bson:"diff_lftsam"`
	DiffVeilig float64 `bson:"diff_veilig"`
}

type CombinedFinancial struct {
	Labels []string `bson:"Labels,omitempty" json:",omitempty"`
	Data   []int64  `bson:"Data,omitempty" json:",omitempty"`
}

type Address struct {
	City           string
	Street         string
	Postcode       Postcode
	Number         string
	NumberOnly     int64
	NumberLetter   string
	NumberAddition string
}

type Postcode struct {
	P4 string
	P6 string
}
