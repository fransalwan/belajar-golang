package main

import "fmt"

func main() {
	// students := []map[string]string{
	// 	{"name": "Alice", "score": "A"},
	// 	{"name": "Bob", "score": "B"},
	// 	{"name": "Charlie", "score": "C"},
	// }

	// for _, student := range students {
	// 	fmt.Println("Name:", student["name"], "Score:", student["score"])
	// }

	scores := [8]int{80, 90, 70, 92, 95, 60, 75, 88}

	// ? soal 1: hitung rata-rata scores
	// var total int

	// for _, score := range scores {
	// 	total += score
	// }
	// fmt.Println("Total nilai:", total)

	// average := float64(total) / float64(len(scores))
	// fmt.Printf("Rata-rata nilai: %.2f\n", average)

	// ? soal 2: buat slice untuk menyimpan nilai yang lebih dari 90
	var goodScore []int

	for _, score := range scores {
		if score > 90 {
			goodScore = append(goodScore, score)
		}
	}

	fmt.Println("Nilai yang lebih dari 90:", goodScore)
}
