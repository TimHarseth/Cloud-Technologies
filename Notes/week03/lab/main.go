package main

import (
	"Notes/week03/lab/mypackage"
	"fmt"
)

func main() {

	// Create a new student
	student1 := mypackage.Student{Age: 20, Name: "John Doe", StudentID: "123456"}
	student2 := mypackage.Student{Age: 21, Name: "Jane Doe", StudentID: "654321"}
	student3 := mypackage.Student{Age: 22, Name: "Alice Smith", StudentID: "112233"}
	student4 := mypackage.Student{Age: 23, Name: "Bob Brown", StudentID: "223344"}
	student5 := mypackage.Student{Age: 24, Name: "Charlie Johnson", StudentID: "334455"}

	// Print the student
	fmt.Println(student1)
	fmt.Println("Using the String() method:")
	fmt.Println(student1.String())

	// Store several students
	sArray := make([]mypackage.Student, 5)
	sArray[0] = student1
	sArray[1] = student2
	sArray[2] = student3
	sArray[3] = student4
	sArray[4] = student5

	// Print all students
	fmt.Println("All Students:")
	for i := 0; i < len(sArray); i++ {
		fmt.Println(sArray[i].String())
	}
}
