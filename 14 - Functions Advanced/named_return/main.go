package main

import "fmt"

func math(num1, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2

	return
}

func main() {
	a, b := math(123, 231)
	fmt.Println(a, b)
}
