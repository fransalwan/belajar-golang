package main

import "fmt"

func main() {
	// ? contoh 1:  penggunaan conditional statement
	// age := 9

	// if age < 18 {
	// 	fmt.Println("You are a minor.")
	// } else {
	// 	fmt.Println("You are an adult.")
	// }

	// ? contoh 2: penggunaan conditional statement
	score := 40

	var grade string
	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score <= 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}

	fmt.Println("Your grade is:", grade)
}
