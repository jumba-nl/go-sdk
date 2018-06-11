package model

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const NOTIF_TYPE_CLAIMED = "claimed-property"
const NOTIF_TYPE_MODIFIED = "modified-property"
const NOTIF_TYPE_QUERY = "modified-query"
const NOTIF_TYPE_NEIGHBOURHOOD = "neighbourhood"

const NOTIF_PROPERTY_MAIL_TMPL = "jumba-property-updated"
const NOTIF_PROPERTIES_MAIL_TMPL = "jumba-properties-updated"
const NOTIF_NEIGHBOURHOOD_MAIL_TMPL = "jumba-neighbourhood-updated"

const NOTIF_CHANGE_PRICE = "De huidige prijs is van € %s naar € %s veranderd."
const NOTIF_CHANGE_FORSALE_TRUE = "De woning staat nu te koop."
const NOTIF_CHANGE_FORSALE_FALSE = "De woning staat niet meer te koop."
const NOTIF_CHANGE_SOLD = "De woning is verkocht onder voorbehoud."
const NOTIF_CHANGE_CLAIMED = "De woning is door een gebruiker geclaimed."
const NOTIF_CHANGE_COMMENT_NEW = "Er is een nieuwe reactie geplaatst over de omgeving door %s."

var CLAIMED_MODIFIED = Modified{Field: "claimed", Old: "unclaimed", New: "claimed"}

type NotifRecip struct {
	ID      bson.ObjectId `bson:"_id"`
	Profile string        `bson:"profile"`
}

type Notification struct {
	URL  string
	Date time.Time
	Type string

	// queries to which the notification is matched
	Queries []struct {
		Profile string        `bson:",omitempty"`
		Query   bson.ObjectId `bson:",omitempty"`
	} `json:"-" bson:",omitempty"`

	// notifications are matched using profile
	Profiles []string `json:"-"`

	// describes changes to a property
	Property NotificationProperty `json:",omitempty"`
	Modified []Modified           `json:",omitempty"`
}

type NotificationLastSeen struct {
	ID      bson.ObjectId `bson:"_id"`
	Profile string
	Seen    time.Time
}

// swagger:model NotificationCount
type NotificationCount struct {
	New   int64
	Since time.Time
}

type NotificationProperty struct {
	Image       string `json:",omitempty" bson:",omitempty"`
	City        string `json:",omitempty" bson:",omitempty"`
	Postalcode  string `json:",omitempty" bson:",omitempty"`
	Postalcode4 string `json:",omitempty" bson:",omitempty"`
	Street      string `json:",omitempty" bson:",omitempty"`
	Number      string `json:",omitempty" bson:",omitempty"`
	Province    string `json:",omitempty" bson:",omitempty"`
	County      string `json:",omitempty" bson:",omitempty"`
}

type Modified struct {
	Field string
	Old   string `json:",omitempty" bson:",omitempty"`
	New   string `json:",omitempty" bson:",omitempty"`
}

type CanBeBson bson.ObjectId

func (s *CanBeBson) SetBSON(raw bson.Raw) error {
	var value string
	if err := raw.Unmarshal(&value); err != nil {
		return err
	} else {
		if bson.IsObjectIdHex(value) {
			*s = CanBeBson(bson.ObjectIdHex(value))
		} else {
			var identifier bson.ObjectId
			if err := raw.Unmarshal(&identifier); err == nil {
				*s = CanBeBson(identifier)
			} else {
				return err
			}
		}
	}

	return nil
}

func GetLastSeenDateNotifications(profile string) (NotificationLastSeen, error) {
	var last NotificationLastSeen
	err := DB().C("notifications_seen").Find(bson.M{"profile": profile}).One(&last)
	return last, err
}

func UpdateLastSeenDateNotification(profile string) error {
	last, err := GetLastSeenDateNotifications(profile)
	if err != nil {
		last.Profile = profile
	}

	if !last.ID.Valid() {
		last.ID = bson.NewObjectId()
	}

	last.Seen = time.Now()
	_, err = DB().C("notifications_seen").UpsertId(last.ID, last)
	return err
}

func GetNotifications(profile string, start time.Time) ([]Notification, error) {
	notifications := []Notification{}
	err := DB().C("notifications").Find(bson.M{"date": bson.M{"$gte": start}, "$or": []bson.M{
		bson.M{"profile": profile},
		// TODO(t) make sure this really needs $in operator, because len(args) is always 1.
		bson.M{"profiles": bson.M{"$in": []string{profile}}},
	}}).Sort("-date").All(&notifications)

	return notifications, err
}

func GetNewNotificationCount(profile string) NotificationCount {
	last, _ := GetLastSeenDateNotifications(profile)
	notifications, _ := GetNotifications(profile, last.Seen)
	return NotificationCount{
		New:   int64(len(notifications)),
		Since: last.Seen,
	}
}

func (c *Notification) Save() error {
	_, err := DB().C("notifications").UpsertId(bson.NewObjectId(), bson.M{"$set": c})
	return err
}

func CreateNeighbourhoodChangesNotification(postalcode string) (Notification, []string) {
	rec := []NotifRecip{}
	// search within favorites
	// notify users with notification / email trigger
	DB().C("favorites").Find(bson.M{"$or": []bson.M{
		bson.M{"property": bson.M{"$regex": bson.RegEx{fmt.Sprintf("^%s", postalcode), ""}}, "active": true},
		bson.M{"url": bson.M{"$regex": bson.RegEx{fmt.Sprintf("^http://jumba.nl/%s", postalcode), ""}}},
	}}).All(&rec)

	p := []string{}
	for _, s := range rec {
		p = append(p, s.Profile)
	}

	return Notification{Profiles: p, Property: NotificationProperty{
		Postalcode: postalcode,
	}}, p
}

func CreatePropertyChangesNotification(hit Search) (Notification, []string) {
	rec := []NotifRecip{}
	// search within favorites
	// notify users with notification / email trigger
	DB().C("favorites").Find(bson.M{"$or": []bson.M{
		bson.M{"property": hit.Payload.ID, "active": true},
		bson.M{"url": fmt.Sprintf("http://jumba.nl/%s", hit.Payload.ID)},
	}}).All(&rec)

	p := []string{}
	for _, s := range rec {
		p = append(p, s.Profile)
	}

	return Notification{URL: fmt.Sprintf("/%s/%s/%s", hit.Payload.Main.City, hit.Payload.Main.Street, hit.Payload.Main.Number), Profiles: p, Date: hit.LastCheck, Property: NotificationProperty{
		Image:      hit.Payload.Image,
		City:       hit.Payload.Main.City,
		Postalcode: hit.Payload.Main.Postcode.P6,
		Street:     hit.Payload.Main.Street,
		Number:     hit.Payload.Main.Number,
		Province:   hit.Payload.Province,
		County:     hit.Payload.County,
	}}, p
}
