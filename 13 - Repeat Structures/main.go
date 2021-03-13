package main

import (
	"fmt"
)

func main() {

	// i := 1

	// fmt.Println("Iniciou o valor com 1")
	// for i < 10 {
	// 	// time.Sleep(time.Second)
	// 	i++
	// 	fmt.Println(i)
	// }

	// fmt.Println("Iniciou o valor de j com 1")
	// for j := 1; j < 10; j += 4 {
	// 	time.Sleep(time.Second)
	// 	fmt.Println(j)
	// }

	// nomes := []string{"João", "Dhaniel", "Davi"}

	// for _, value := range nomes {
	// 	fmt.Println(value)
	// }

	// for _, value := range "ABCDE" {
	// 	fmt.Println(string(value)) // Vem por padrão os códigos da tabela ASC
	// }

	// user := map[string]string{
	// 	"name": "Dhaniel",
	// }

	// for key, value := range user {
	// 	fmt.Println(key, value)
	// }

	for i := 1; true; i++ {
		fmt.Println(i)
		if i == 10 {
			break
		}
	}
}
