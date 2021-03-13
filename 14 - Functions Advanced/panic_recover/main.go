package main

import (
	"fmt"
	"time"
)

func recoverExec() {
	fmt.Println("Tentando recuperar a execução...")
	time.Sleep(time.Second * 3)
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func studentIsApproved(num1, num2 float32) bool {
	defer recoverExec()
	media := (num1 + num2) / 2

	time.Sleep(time.Second)
	fmt.Println("Calculando média...")
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	fmt.Println("Média calculada")
	time.Sleep(time.Second)

	panic("A MÉDIA É EXATAMENTE 6")
}

func main() {
	fmt.Println(studentIsApproved(12, 6))
	fmt.Println("Pós execução")
}
