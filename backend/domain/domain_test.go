package domain

import "fmt"

func assertCourse(actual, expected *Course) error {
	if actual.Name != expected.Name {
		return reportUnexpected("name", actual.Name, expected.Name)
	}
	if err := assertCoursePrice(actual.Price, expected.Price); err != nil {
		return fmt.Errorf("unexpected price: %w", err)
	}

	return nil
}

func assertCoursePrice(actual, expected *CoursePrice) error {
	if actual.Amount != expected.Amount {
		return reportUnexpected("amount", actual.Amount, expected.Amount)
	}
	if actual.Currency != expected.Currency {
		return reportUnexpected("currency", actual.Currency, expected.Currency)
	}

	return nil
}

func reportUnexpected(field string, actual, expected interface{}) error {
	return fmt.Errorf("unexpected %s: got %v, expect %v", field, actual, expected)
}
