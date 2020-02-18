package domain

import (
	"net/url"
	"time"
)

type Consultation struct {
	Client     *User
	Consultant *Consultant
	DateTime   *ConsultationDateTime
	Location   *ConsultationLocation
}

type ConsultationDateTime struct {
	From, To time.Time
}

type ConsultationLocation struct {
	Name string
	URI  url.URL
}
