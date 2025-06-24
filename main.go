package main

import (
	"fmt"
	"golang-basic/calculation"
)

func main() {
	fmt.Println("Hello, World!")
	result := calculation.Add(5, 3)
	fmt.Printf("The result of adding 5 and 3 is: %d\n", result)
}
