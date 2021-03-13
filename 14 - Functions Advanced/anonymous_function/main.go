package main

import "fmt"

func main() {

	func(text string) {
		fmt.Println(text)
	}("passando parametros")

	function := func(text string) {
		fmt.Println(text)
	}

	function("Oi")
}
