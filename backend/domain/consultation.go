package domain

import (
	"net/url"
	"time"
)

type Consultation struct {
	Client   *User
	Course   *Course
	Content  string
	DateTime *ConsultationDateTime
	Location *ConsultationLocation
}

type ConsultationDateTime struct {
	From, To time.Time
}

type ConsultationLocation struct {
	Name string
	URI  url.URL
}
