package main

import "fmt"

func main() {
	// ARITIMÉTICOS
	var soma int8 = 1 + 2
	var sub int8 = 1 - 2
	var div int8 = 1 / 2
	var mult int8 = 1 * 2
	var modulo int8 = 1 % 2

	fmt.Println(soma)
	fmt.Println(sub)
	fmt.Println(div)
	fmt.Println(mult)
	fmt.Println(modulo)

	// RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// LÓGICOS
	fmt.Println("------------------")
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)

	// UNÁRIOS
	fmt.Println("------------------")
	num := 10
	num++

	num2 := 10
	num2 += 15

	num3 := 10
	num3--

	num4 := 10
	num4 -= 15

	num5 := 10
	num5 *= 15

	num6 := 10
	num6 /= 15

	num7 := 10
	num7 %= 15

	fmt.Println(num)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
	fmt.Println(num5)
	fmt.Println(num6)
	fmt.Println(num7)
}
