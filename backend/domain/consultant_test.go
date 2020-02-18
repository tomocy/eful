package domain

import (
	"testing"

	"github.com/tomocy/eful/backend/domain/currency"
)

func TestConsultant_CreateCourse(t *testing.T) {
	con := &Consultant{
		Name:  "name",
		Title: "title",
	}
	cou := &Course{
		Name: "name",
		Price: &CoursePrice{
			Amount:   1000,
			Currency: currency.JPY,
		},
	}

	tests := map[string]func(t *testing.T){
		"success": func(t *testing.T) {
			con.Courses = nil
			if err := con.CreateCourse(cou); err != nil {
				t.Errorf("should have created course: %w", err)
				return
			}
			if len(con.Courses) != 1 {
				t.Errorf("should have had the created course: %w", reportUnexpected("len of courses", len(con.Courses), 1))
				return
			}
			if err := assertCourse(con.Courses[0], cou); err != nil {
				t.Errorf("should have had expected course: unexpected course: %w", err)
				return
			}
		},
		"failed due to duplicated courses": func(t *testing.T) {
			con.Courses = []*Course{cou}
			if err := con.CreateCourse(cou); err == nil {
				t.Error("should have failed to create course due to duplicated courses")
				return
			}
			if len(con.Courses) != 1 {
				t.Errorf("should not have had the created course: %w", reportUnexpected("len of courses", len(con.Courses), 0))
				return
			}
			if err := assertCourse(con.Courses[0], cou); err != nil {
				t.Errorf("should not have affect other courses: unexpected course: %w", err)
				return
			}
		},
	}

	for name, test := range tests {
		t.Run(name, test)
	}
}
