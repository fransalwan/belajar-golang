package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

// menggunakan pointer pada method
func (s *Student) graduate() {
	s.Name = s.Name + " S.Kom"
	// fmt.Println("Inside graduate function:", student.Name)
}

func main() {
	student := Student{1, "John Doe", 3.5}

	fmt.Println(student.Name)

	student.graduate()

	fmt.Println(student.Name) // Name remains unchanged
}

// menggunkan pointer pada function
// func graduate(student *Student) {
// 	student.Name = student.Name + " S.Kom"
// 	// fmt.Println("Inside graduate function:", student.Name)
// }
