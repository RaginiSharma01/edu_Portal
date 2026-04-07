package models

import "time"

type CreateEvent struct {
	Title       string `json:"title"`
	EventDate   string `json:"eventDate"`
	Venue       string `json:"venue"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	EventDate   time.Time `json:"eventDate"`
	Venue       string    `json:"venue"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
}
