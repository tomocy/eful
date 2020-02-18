package domain

import (
	"github.com/tomocy/eful/backend/domain/currency"
)

type Course struct {
	Consultant *Consultant
	Name       string
	Price      *CoursePrice
}

type CoursePrice struct {
	Amount   float32
	Currency currency.Currency
}
