package main

import "fmt"

func sum(num1 int8, num2 int8) int8 {
	return num1 + num2
}

func math(num1, num2 int8) (int8, int8) {
	soma := num1 + num2
	sub := num1 - num2

	return soma, sub
}

func main() {
	soma := sum(10, 20)

	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	f("Texto agora")

	sum, _ := math(10, 14)
	fmt.Println(sum)

}
