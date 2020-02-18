package domain

import "fmt"

type Consultant struct {
	Name    string
	Title   string
	Courses []*Course
}

func (c *Consultant) CreateCourse(co *Course) error {
	for _, stored := range c.Courses {
		if stored.Name == co.Name {
			return fmt.Errorf("duplicated course")
		}
	}

	c.Courses = append(c.Courses, co)

	return nil
}
