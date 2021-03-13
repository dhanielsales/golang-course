package main

import (
	"fmt"
)

func func1() {
	fmt.Println("func1")
}

func func2() {
	fmt.Println("func2")
}

func studentIsApproved(studant string, grades ...float64) {

	total := 0.0

	for _, value := range grades {
		total += value
	}

	media := total / float64(len(grades))

	fmt.Println("Nota calculada, o aluno, média =>", studant, media)
}

func main() {
	defer func1() // Por ultimo

	studentIsApproved("Dhaniel", 9, 5, 8, 4, 5, 6, 7, 8, 9.10)

	func2()
}
