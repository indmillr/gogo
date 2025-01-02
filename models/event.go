package models

import "time"

type Event struct {
	ID int
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int
}

var events = []Event{}

// ----- SAVE
func (e Event) Save() {
	// --- TODO: add to DB
	events = append(events, e)
}

// ----- Get all available Events
func GetAllEvents() []Event {
	return events
}