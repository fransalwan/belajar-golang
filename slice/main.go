package main

import "fmt"

func main() {
	var gamingConsoles []string

	gamingConsoles = append(gamingConsoles, "PlayStation 1")
	gamingConsoles = append(gamingConsoles, "PlayStation 2")
	gamingConsoles = append(gamingConsoles, "PlayStation 3")

	// fmt.Println("Gaming Consoles:", gamingConsoles)
	for _, console := range gamingConsoles {
		fmt.Println("Gaming Console:", console)
	}
}
