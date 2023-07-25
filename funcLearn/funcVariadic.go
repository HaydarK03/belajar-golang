package main

// function yg memiliki variable argumen disebut variadic func
// varags datanya dapat menerima lebih dari satu input
func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}
