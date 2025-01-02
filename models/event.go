package models

import (
	"time"

	"example.com/gogo/db"
)

type Event struct {
	ID int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int
}

var events = []Event{}

// ----- SAVE
func (e Event) Save() error {
	// --- Question Marks are used in query strings to help with sanitation
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	
	// --- defer Close of stmt to end of function
	defer stmt.Close()

	// --- Execute stmt using 'query' where Question Marks are replaced by values passed to stmt.Exec() below
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	// --- Get the ID of the most recently inserted data
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

// ----- Get all available Events
func GetAllEvents() []Event {
	return events
}