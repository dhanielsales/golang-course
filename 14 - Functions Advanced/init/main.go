package main

import "fmt"

func init() {
	fmt.Println("init") // Sempre será executada primeiro no arquivo - Pode ter uma por arquivo
}

func main() {
	fmt.Println("Main")
}
