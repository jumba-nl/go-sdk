package model

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// swagger:model Profile
type Profile struct {
	ID           string `json:"id" bson:"_id"`
	EmailAddress string `json:"emailaddress,omitempty"`
	Firstname    string `json:"firstname,omitempty"`
	Infix        string `json:"infix,omitempty"`
	Lastname     string `json:"lastname,omitempty"`
	Gender       string `json:"gender,omitempty"`
	Picture      string `json:"picture,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Address      string `json:"address,omitempty"`

	JumbaFtu     bool `json:"jumba_ftu"`
	OptIn        bool `json:"opt_in" bson:"opt_in"`
	EnableOffers bool `json:"enable_offers" bson:"enable_offers"`

	ClaimedFtu       bool      `json:"claimed_ftu" bson:"claimed_ftu"`
	ClaimedProcessed bool      `json:"claimed_processed" bson:"claimed_processed"`
	ClaimedStatus    int       `json:"claimed_status" bson:"claimed_status"`
	ClaimedAddress   string    `json:"claimed_address" bson:"claimed_address"`
	ClaimedDatetime  time.Time `json:"claimed_datetime" bson:"claimed_datetime"`

	Registered time.Time `json:"registered,omitempty"`
	Created    time.Time `json:"created,omitempty"`
	Updated    time.Time `json:"updated,omitempty"`
}

type ProfileSimple struct {
	ID string `json:"id,omitempty" bson:"_id"`

	EmailAddress string `json:"emailaddress,omitempty"`
	Firstname    string `json:"firstname,omitempty"`
	Infix        string `json:"infix,omitempty"`
	Lastname     string `json:"lastname,omitempty"`
	Picture      string `json:"picture,omitempty"`
}

// ProfileAddressClaim represents a claim of an address
// swagger:model ProfileAddressClaim
type ProfileAddressClaim struct {
	ClaimedProcessed bool      `json:"claimed_processed"` // if claim has been accepted
	ClaimedAddress   string    `json:"claimed_address"`   // Postalcode + number (+ additions)
	ClaimedStatus    int       `json:"claimed_status"`    // special status added by admin
	ClaimedDatetime  time.Time `json:"claimed_datetime"`  // time of claim
}

// Enum of possible claimed statusses
const (
	ClaimStatusUnknown      = iota // = no special process currently running
	ClaimStatusIsRental            // = claimed address is a rental property
	ClaimStatusNameMismatch        // = registered Kadaster name doesn't match profile name
	ClaimStatusNameUnclear         // = name is unclear (for example names that are clearly fake: Test Test)
)

func NewProfile() *Profile {
	return &Profile{}
}

// GetJumbaProfile returns a profile from a ID and returns a non nil error value when something went wrong
func GetJumbaProfileByID(i string) (*Profile, error) {
	profile := &Profile{}

	err := DB().C("profile_users").FindId(i).One(profile)

	return profile, err
}

func GetJumbaProfile(q bson.M) (*Profile, error) {
	profile := &Profile{}

	err := DB().C("profile_users").Find(q).One(profile)

	return profile, err
}

func FilterJumbaProfilesSimple(q bson.M) []ProfileSimple {
	var item ProfileSimple = ProfileSimple{}
	list := []ProfileSimple{}

	find := DB().C("profile_users").Find(q)

	items := find.Iter()
	for items.Next(&item) {
		list = append(list, item)
	}

	return list
}

func GetJumbaProfiles(q bson.M) []Profile {
	var item Profile = Profile{}
	list := []Profile{}

	find := DB().C("profile_users").Find(q)

	items := find.Iter()
	for items.Next(&item) {
		list = append(list, item)
	}

	return list
}

// GetProfileSimple returns a simple profile from a ID and returns a non nil error value when something went wrong
func GetJumbaProfileSimple(i string) (*ProfileSimple, error) {
	profile := &ProfileSimple{}

	err := DB().C("profile_users").FindId(i).One(profile)

	return profile, err
}

func GetJumbaProfilesSimple(q bson.M) []ProfileSimple {
	var item ProfileSimple = ProfileSimple{}
	list := []ProfileSimple{}

	find := DB().C("profile_users").Find(q)

	items := find.Iter()
	for items.Next(&item) {
		list = append(list, item)
	}

	return list
}

func SelectFieldsJumbaProfile(q bson.M, fields []string) (*Profile, error) {
	var fieldList bson.M = bson.M{}

	for _, val := range fields {
		fieldList[val] = 1
	}

	profile := &Profile{}

	err := DB().C("profile_users").Find(q).Select(fieldList).One(profile)

	return profile, err
}

func SelectFieldsJumbaProfiles(q bson.M, fields []string) ([]Profile, error) {
	var fieldList bson.M = bson.M{}

	for _, val := range fields {
		fieldList[val] = 1
	}

	profiles := []Profile{}

	err := DB().C("profile_users").Find(q).Select(fieldList).All(&profiles)

	return profiles, err
}

// SaveJumba saves the current profile in this model to the db and returns a non nil error value when something went wrong
func (c *Profile) SaveJumba() error {
	if c.ID == "" {
		return errors.New("invalid profile id")
	}

	_, err := DB().C("profile_users").UpsertId(c.ID, bson.M{"$set": c})
	return err
}

func (p *Profile) Name() string {
	n := p.Firstname + " "
	if p.Infix != "" {
		n += p.Infix + " "
	}

	return n + p.Lastname
}

// AddressClaim returns the address claimed by Profile
func (p *Profile) AddressClaim() ProfileAddressClaim {
	if p != nil {
		return ProfileAddressClaim{
			ClaimedAddress:   p.ClaimedAddress,
			ClaimedDatetime:  p.ClaimedDatetime,
			ClaimedProcessed: p.ClaimedProcessed,
			ClaimedStatus:    p.ClaimedStatus,
		}
	}
	return ProfileAddressClaim{}
}

// ClaimAddress applies the given claim to a Profile
func (p *Profile) ClaimAddress(claim ProfileAddressClaim) {
	p.ClaimedAddress = claim.ClaimedAddress
	p.ClaimedDatetime = claim.ClaimedDatetime
	p.ClaimedProcessed = claim.ClaimedProcessed
	p.ClaimedStatus = claim.ClaimedStatus
}

// Make sure given profile exists as a Jumba profile.
func EnsureProfileExists(p *Profile) error {
	// first time user flag
	p.JumbaFtu = true
	// set created and updated
	now := time.Now()
	p.Created = now
	p.Updated = now

	n, err := DB().C("profile_users").FindId(p.ID).Count()
	if err != nil {
		return err
	}
	if n == 0 {
		err = DB().C("profile_users").Insert(p)
	}
	return err
}
