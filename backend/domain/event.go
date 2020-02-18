package domain

import (
	"net/url"
	"time"
)

type Event struct {
	Client     *User
	Specialist *Specialist
	DateTime   *EventDateTime
	Location   *EventLocation
}

type EventDateTime struct {
	From, To time.Time
}

type EventLocation struct {
	Name string
	URI  url.URL
}
