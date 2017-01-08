package events

import (
	"errors"
	"fmt"
	"reflect"
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

// GetEventNode . . . get an event node
func GetEventNode(identifier string) (Event, error) {
	var event Event

	stmt := `
		MATCH (event:Event)
		WHERE event.UniqueID = {UniqueID}
		RETURN event
	`
	params := neoism.Props{
		"UniqueID": identifier,
	}

	// query results
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
		return event, err
	}

	err = event.FillStruct(res[0].Event.Data)
	if err != nil {
		return event, err
	}
	return event, nil
}

//TODO deleteEvent()

//TODO updateEvent()

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

func (s *Event) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
