package main

import (
	"Notes/week03/lab/mypackage"
	"fmt"
)

func main() {

	// Create a new student
	student := mypackage.Student{Age: 20, Name: "John Doe", StudentID: "123456"}

	// Print the student
	fmt.Println(student)
}
