package models

import (
	"example.com/rest-api/db"
)

type Event struct {
	ID          int64    `json:"id"`
	Name        string `json:"name" bind:"required"`
	Description string `json:"description" bind:"required"`
	DateTime        string `json:"date" bind:"required"`
	Location    string `json:"location" bind:"required"`
	UserID      int    `json:"user_id"`
}

func (e *Event) Save() error {
	// Implementation for saving the event

	query := `INSERT INTO events (name, description, location, dateTime, user_id) VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	e.ID = id

	return nil		
}

func  GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err}

		events = append(events, event)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}	