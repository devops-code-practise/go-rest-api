package models

type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name" bind:"required"`
	Description string `json:"description" bind:"required"`
	DateTime        string `json:"date" bind:"required"`
	Location    string `json:"location" bind:"required"`
	UserID      int    `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() error {
	// Implementation for saving the event
	events = append(events, *e)
	return nil
}

func  GetAllEvents() []Event {
	return events
}	