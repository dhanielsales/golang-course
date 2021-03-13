package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func mudarDados(dados map[string]interface{}) {
	for key, value := range dados {
		fmt.Println(key, value)
	}
}

func main() {
	// generic("aaaa")

	object := map[string]interface{}{
		"name": "Dhaniel",
		"age":  26,
	}

	mudarDados(object)

	// fmt.Println(object)
}
