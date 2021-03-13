package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}

func write(text string, numbers ...int) {
	for _, value := range numbers {
		fmt.Println(text, value)
	}
}

func main() {
	fmt.Println(sum(123, 2, 13, 54, 21, 467, 02, 0))
	write("Texto", 1, 2, 3453, 23, 1123, 6, 4)
}
