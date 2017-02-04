package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Quote struct represents a quote entity
type Quote struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Author string        `json:"author" bson:"author"`
	Quote  string        `json:"quote" bson:"quote"`
	Link   string        `json:"link" bson:"link"`
}
