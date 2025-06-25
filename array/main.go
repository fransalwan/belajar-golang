package main

import "fmt"

func main() {

	language := [...]string{"Go", "Java", "Python", "C++", "JavaScript"}

	// 	fmt.Println("Languages:", language)
	// 	length := len(language)
	// 	fmt.Println("Number of languages:", length)

	for index, lang := range language {
		fmt.Printf("Index: %d, Language: %s\n", index, lang)
	}
}
