package main

import "fmt"

func invertSigned(num int) int {
	return num * -1
}

func invertSignedWithPointer(num *int) int {
	*num = *num * -1

	return *num
}

func main() {
	num := 10
	fmt.Println(num)

	num2 := invertSigned(num)

	fmt.Println(num)
	fmt.Println(num2)

	fmt.Println("-------------------------")

	num3 := 10
	fmt.Println(num3)

	num4 := invertSignedWithPointer(&num3)

	fmt.Println(num3)
	fmt.Println(num4)
}
