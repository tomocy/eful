package domain

import (
	"net/url"
	"time"
)

type Event struct {
	Client     *User
	Consultant *Consultant
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
