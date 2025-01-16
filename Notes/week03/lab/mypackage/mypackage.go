package mypackage

import "strconv"

type Student struct {
	Age       int
	Name      string
	StudentID string
}

func (s *Student) String() string {
	return "Student name: " + s.Name + " age: " + strconv.Itoa(s.Age) + " studentID: " + s.StudentID
}
