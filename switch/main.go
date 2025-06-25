package main

import "fmt"

func main() {
	number := 2

	switch number {
	case 1:
		fmt.Println("Number is one")
	case 2:
		fmt.Println("Number is two")
	case 3:
		fmt.Println("Number is three")
	default:
		fmt.Println("Number is something else")
	}
	// This is the main function where the program starts.
	// You can add your code here to call functions from the calculation package.
	// For example, you can call Add or Multiply functions.
}
