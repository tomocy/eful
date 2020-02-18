package domain

import "fmt"

func reportUnexpected(field string, actual, expected interface{}) error {
	return fmt.Errorf("unexpected %s: got %v, expect %v", field, actual, expected)
}
