package models

import (
	"example.com/rest-api/db"
)

type Event struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" bind:"required"`
	Description string `json:"description" bind:"required"`
	DateTime    string `json:"date" bind:"required"`
	Location    string `json:"location" bind:"required"`
	UserID      int64  `json:"user_id"`
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


func GetEventByID(id int64) (*Event, error) {
	var event Event
	query := `SELECT * FROM events WHERE id = ?`
	err := db.DB.QueryRow(query, id).Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
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


func Update(event *Event) error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, dateTime = ? WHERE id = ?`
	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.ID)

	if err != nil {
		return err
	}

	return nil
}

func DeleteEvent(id int64) error {
	query := `DELETE FROM events WHERE id = ?`
	_, err := db.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}