package mypackage

import "strconv"

type Student struct {
	Age       int
	Name      string
	StudentID string
}

func (s *Student) String() string {
	return "Student with name " + s.Name + " and age " + strconv.Itoa(s.Age)
}
