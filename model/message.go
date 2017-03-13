package model

import "time"

// Message is a message passed into the system.
type Message struct {
	TimestampUTC time.Time
	UUID         string
	Body         string
}

// TableName returns the table name.
func (m Message) TableName() string {
	return "message"
}
