package events

import (
	"time"

	uuid "github.com/satori/go.uuid"

	neoism "gopkg.in/jmcvetta/neoism.v1"
)

var (
	Db *neoism.Database
)

func init() {
	var err error
	Db, err = neoism.Connect("http://neo4j:password@localhost:7474/db/data")
	if err != nil {
		panic(err)
	}
}

// Event ... event struct for neo4j event nodes
type Event struct {
	Properties properties
}

// An events properties
type properties struct {
	Name        string    `json:"name"`
	DateCreated time.Time `json:"date"`
	Description string    `json:"description"`
	Keywords    []string  `json:"keywords"`
	TypeOfEvent string    `json:"typeofevent"`
	Emblem      string    `json:"emblem"`
	Rating      float64   `json:"rating"`
	UniqueID    string
	Location    location
}

type location struct {
	StreetAddress string `json:"address"`
	City          string `json:"city"`
	State         string `json:"state"`
	ZipCode       string `json:"zipcode"`
}

// EventRelationships ... neo4j relationships associated with Event nodes
var EventRelationships = map[string]interface{}{
	"Attended":      "ATTENDED",
	"IsAttending":   "IS_ATTENDING",
	"PostedVideo":   "POSTED_A_VIDEO",
	"PostedComment": "POSTED_A_COMMENT",
	"PostedPhoto":   "POSTED_A_PHOTO",
	"IsLocated":     "IS_LOCATED",
}

//liveProperties ... properties that contain real time information or data
//such as photos, videos, and comments
type liveProperties struct {
	Photo    photo
	Comments []comment
}

type photo struct {
	URI        string    `json:"uri"`
	DatePosted time.Time `json:"date"`
	UserID     string    `json:"userID"`
}

type comment struct {
	Message    string    `json:"message"`
	DatePosted time.Time `json:"date"`
	UserID     string    `json:"userID"`
}

// LivePropertyRelationships ... neo4j relationships associated with live property nodes
var LivePropertyRelationships = map[string]interface{}{
	"AssociatedWith": "ASSOCIATED_WITH",
}

// CreateEventNode . . . create a new event node from Event struct
func CreateEventNode(event Event) (Event, error) {
	uid := uuid.NewV4().String()
	node, err := Db.CreateNode(neoism.Props{
		"name":           event.Properties.Name,
		"date":           event.Properties.DateCreated,
		"description":    event.Properties.Description,
		"keywords":       event.Properties.Keywords,
		"type-of-event":  event.Properties.TypeOfEvent,
		"emblem":         event.Properties.Emblem,
		"rating":         event.Properties.Rating,
		"street-address": event.Properties.Location.StreetAddress,
		"city":           event.Properties.Location.City,
		"state":          event.Properties.Location.State,
		"zip-code":       event.Properties.Location.ZipCode,
		"unique-id":      uid,
	})
	if err != nil {
		return event, err
	}

	node.AddLabel("Event")
	event.Properties.UniqueID = uid
	return event, nil
}

//TODO getEvent()

//TODO deleteEvent()

//TODO updateEvent()
