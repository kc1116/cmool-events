package events

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"errors"

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
	Properties Properties
}

// An events properties
type Properties struct {
	Name          string    `json:"Name"`
	DateCreated   time.Time `json:"Date"`
	Description   string    `json:"Description"`
	Keywords      []string  `json:"Keywords"`
	TypeOfEvent   string    `json:"TypeOfEvent"`
	Emblem        string    `json:"Emblem"`
	Rating        float64   `json:"Rating"`
	StreetAddress string    `json:"StreetAddress"`
	City          string    `json:"City"`
	State         string    `json:"State"`
	Zipcode       string    `json:"ZipCode"`
	UniqueID      string    `json:"UniqueID"`
	Location      location
}

type location struct {
	StreetAddress string `json:"Address"`
	City          string `json:"City"`
	State         string `json:"State"`
	ZipCode       string `json:"Zipcode"`
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
	URI        string    `json:"Uri"`
	DatePosted time.Time `json:"Date"`
	UserID     string    `json:"UserID"`
}

type comment struct {
	Message    string    `json:"Message"`
	DatePosted time.Time `json:"Date"`
	UserID     string    `json:"UserID"`
}

// LivePropertyRelationships ... neo4j relationships associated with live property nodes
var LivePropertyRelationships = map[string]interface{}{
	"AssociatedWith": "ASSOCIATED_WITH",
}

// CreateEventNode . . . create a new event node from Event struct
func CreateEventNode(event Event) (Event, error) {
	uid := uuid.NewV4().String()
	node, err := Db.CreateNode(neoism.Props{
		"Name":          event.Properties.Name,
		"DateCreated":   event.Properties.DateCreated,
		"Description":   event.Properties.Description,
		"Keywords":      event.Properties.Keywords,
		"TypeOfEvent":   event.Properties.TypeOfEvent,
		"Emblem":        event.Properties.Emblem,
		"Rating":        event.Properties.Rating,
		"StreetAddress": event.Properties.Location.StreetAddress,
		"City":          event.Properties.Location.City,
		"State":         event.Properties.Location.State,
		"Zipcode":       event.Properties.Location.ZipCode,
		"UniqueID":      uid,
	})
	if err != nil {
		return event, err
	}

	node.AddLabel("Event")
	event.Properties.UniqueID = uid
	return event, nil
}

// GetEventNode . . . get an event node. returns properties assiciated with that node
func GetEventNode(identifier string) (map[string]interface{}, error) {
	stmt := `
		MATCH (event:Event)
		WHERE event.UniqueID = {uid}
		RETURN event
	`
	params := neoism.Props{
		"uid": identifier,
	}

	// query results
	//res := []Properties{}

	res := []struct {
		Event neoism.Node
	}{}

	cq := neoism.CypherQuery{
		Statement:  stmt,
		Parameters: params,
		Result:     &res,
	}

	err := Db.Cypher(&cq)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		err := errors.New("Event node not found.")
		return nil, err
	}

	return res[0].Event.Data, nil
}

//TODO deleteEvent()

//TODO updateEvent()
