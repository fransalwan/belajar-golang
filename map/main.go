package main

import "fmt"

func main() {
	// var myMap map[string]int

	// myMap = map[string]int{}

	// myMap["ruby"] = 9
	// myMap["python"]x = 10
	// myMap["golang"] = 9

	// fmt.Println("Programming Languages and their ratings:", myMap)

	myMap := map[string]string{
		"ruby":       "is beautiful",
		"golang":     "is fast",
		"javascript": "is versatile",
	}

	// for index, value := range myMap {
	// 	fmt.Println("Programming Language:", index, "Description:", value)
	// }

	// for key, value := range myMap {
	// 	fmt.Printf("Programming Language: %s, Description: %s\n", key, value)
	// }

	// fmt.Println("=============")

	// delete(myMap, "ruby")

	// for key, value := range myMap {
	// 	fmt.Printf("Programming Language: %s, Description: %s\n", key, value)
	// }

	// cek value in myMap variable

	if value, ok := myMap["golang"]; ok {
		fmt.Println("Golang is present in the map with description:", value)
	} else {
		fmt.Println("Golang is not present in the map")
	}
}
