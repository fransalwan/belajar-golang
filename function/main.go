package main

import "fmt"

func main() {
	// printMyResult("Hello, World!")
	// printMyResult("This is a Go function that prints a sentence.")
	// printMyResult("You can call this function with any string you want to print.")

	// result := add(5, 10)
	// fmt.Println(result)

	// luas, keliling := calculate(10, 5)
	// println("Luas:", luas)
	// println("Keliling:", keliling)

	// ? quiz 1

	// scores := []int{90, 85, 78, 92, 88}
	// total := sum(scores)
	// println("Total score:", total)

	// ? quiz 2

	// result, err := calculate(10, 2, "+")
	// result, err := calculate(10, 2, "-")
	// result, err := calculate(10, 2, "*")
	result, err := calculate(10, 2, "/")
	// result, err := calculate(10, 2, "=")

	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Result:", result)
	}
}

func calculate(a, b int, operator string) (result int, errorResult error) {

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		errorResult = fmt.Errorf("unknown operator: %s", operator)
	}

	return result, errorResult
}

// case "-":
// 	return a - b, nil
// case "*":
// 	return a * b, nil
// case "/":
// 	if b == 0 {
// 		return 0, fmt.Errorf("division by zero")
// 	}
// 	return a / b, nil
// default:
// 	return 0, fmt.Errorf("unknown operator: %s", operator)

// func sum(numbers []int) int {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
// 	return total
// }

// contoh predefined function
// func calculate(panjang, lebar int) (luas int, keliling int) {
// 	luas = panjang * lebar
// 	keliling = 2 * (panjang + lebar)

// 	return luas, keliling
// }

// func calculate(panjang, lebar int) (int, int) {
// 	luas := panjang * lebar
// 	keliling := 2 * (panjang + lebar)

// 	return luas, keliling
// }

// func add(a, b int) int {
// 	return a + b
// }

// func printMyResult(sentence string) {
// 	fmt.Println(sentence)
// }
