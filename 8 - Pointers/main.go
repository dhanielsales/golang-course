package main

import "fmt"

func main() {
	fmt.Println("OI")

	var var1 int8 = 10
	var var2 int8 = var1

	fmt.Println(var1, var2)

	var1++

	fmt.Println(var1, var2)

	var var3 int8 = 100
	var var4 *int8 = &var3

	fmt.Println(var3, var4)
	fmt.Println(*var4) // Desreferenciação => Dereferencing
	var3++
	fmt.Println(*var4) // Desreferenciação => Dereferencing
}
