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
	UserID int64
}

var events = []Event{}

// ----- SAVE
func (e *Event) Save() error {
	// --- Question Marks are used in query strings to help with sanitation
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	
	// ----- Prepare is Optional. It is best used when a statement needs to be executed multiple times, but only if is has not been .Close() first
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
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	// --- use Query instead of Exec. Commonly 'Query' is used when just fetching rows, 'Exec' is used when items are changing (INSERT)
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	// --- 'Next' returns True if there are rows left, False if not
	for rows.Next() {
		var event Event 
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := `
	INSERT INTO registrations(event_id, user_id) VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}