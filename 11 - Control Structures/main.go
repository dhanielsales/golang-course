package main

import "fmt"

func main() {
	num := 10

	if num > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if num2 := num; num2 > 0 {
		fmt.Println("Maior que 0")
	}

	if num3 := num; num3 < 0 {
		fmt.Println("Menor que 0")
	} else if num3 > 10 {
		fmt.Println("Maior que 10")
	} else {
		fmt.Println("Entre 0 e 10")
	}
}
