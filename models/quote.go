package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Quote struct represents a quote entity
type Quote struct {
	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Author string        `json:"quoteAuthor" bson:"author"`
	Quote  string        `json:"quoteText" bson:"quote"`
	Link   string        `json:"quoteLink" bson:"link"`
}
