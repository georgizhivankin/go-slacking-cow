package models

import "github.com/satori/go.uuid"

// Quote struct represents a quote entity
type Quote struct {
	ID     uuid.UUID `db:"id" json:"id" valid:"-"`
	Author string    `db:"author" json:"author" valid:"optional"`
	Quote  string    `db:"quote" json:"quote" valid:"required"`
	Link   string    `db:"link" json:"link" valid:"optional,url"`
}
