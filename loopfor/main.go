package main

import (
	"fmt"
)

func main() {

	// ? for loop
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("Loop iteration:", i)
	// }

	// ? while loop
	// i := 1
	// for i <= 100 {
	// 	fmt.Println("Loop iteration:", i)
	// 	i++
	// }

	// ? foreach loop
	title := "Golang the best language"

	// for _, letter := range title {
	// 	// fmt.Printf("Index: %d, Letter: %c\n", index, letter)
	// 	fmt.Println("Letter:", string(letter))
	// }

	// loop just a, i, u, e, o in title variable value
	for index, letter := range title {
		letterString := string(letter)

		// if letter == 'a' || letter == 'i' || letter == 'u' || letter == 'e' || letter == 'o' {
		// 	fmt.Println("index:", index, "Letter:", string(letter))
		// }

		switch letterString {

		case "a", "i", "u", "e", "o":
			fmt.Println("index:", index, "Letter:", letterString)
		}
	}

}
