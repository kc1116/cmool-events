package events

import (
	"time"
)

// Event ... event struct for neo4j event nodes
type Event struct {
	Properties properties
}

// An events properties
type properties struct {
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Keywords    []string  `json:"keywords"`
	TypeOfEvent string    `json:"typeofevent"`
	Emblem      string    `json:"emblem"`
	Rating      string    `json:"rating"`
}

// EventRelationships ... neo4j relationships associated with Event nodes
var EventRelationships = map[string]interface{}{
	"Attended":      "ATTENDED",
	"IsAttending":   "IS_ATTENDING",
	"PostedVideo":   "POSTED_A_VIDEO",
	"PostedComment": "POSTED_A_COMMENT",
	"PostedPhoto":   "POSTED_A_PHOTO",
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
