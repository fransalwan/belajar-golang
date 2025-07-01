package main

import "fmt"

func main() {
	// // Create a variable of type int
	// var x int = 42

	// // Create a pointer to the variable x

	// var p *int = &x

	// //  the memory address of x
	// println("Address of x:", p)

	// // Print the value of x using the pointer
	// println("Value of x using pointer:", *p)

	// // Change the value of x using the pointer
	// *p = 100

	// // Print the new value of x
	// println("New value of x:", x)

	number := 5
	fmt.Println("Alamat memory awal :", &number)
	fmt.Println("Nilai awal :", number)

	change(&number, 100)

	fmt.Println("Nilai akhir :", number)
	fmt.Println("Alamat memory akhir :", &number)
}

func change(old *int, new int) {
	*old = new
	fmt.Println("Nilai baru :", old)
	fmt.Println("di dalam function: ", *old)
}
