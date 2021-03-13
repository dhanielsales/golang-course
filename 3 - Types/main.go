package main

import (
	"errors"
	"fmt"
)

func main() {
	//////////////////// Numeros inteiros positivos e negativos
	var num int8 = 100
	fmt.Println(num)

	var num2 int16 = -10000
	fmt.Println(num2)

	var num3 int32 = 1000000000
	fmt.Println(num3)

	var num4 int64 = -1000000000000000000
	fmt.Println(num4)

	var num5 int = 1000000000000000000 // Usa a arquitetura do computador 32/64
	fmt.Println(num5)

	// Numeros inteiros positivos
	var unum uint8 = 100
	fmt.Println(unum)

	var unum2 uint16 = 10000
	fmt.Println(unum2)

	var unum3 uint32 = 1000000000
	fmt.Println(unum3)

	var unum4 uint64 = 1000000000000000000
	fmt.Println(unum4)

	var unum5 uint = 1000000000000000000 // Usa a arquitetura do computador 32/64
	fmt.Println(unum5)

	// Alias
	// INT32 = RUNE => Numeros que representam caracteres principalmente da tabela ASC
	var runevar rune = 1000000000
	fmt.Println(runevar)

	// UINT8 = BYTE
	var bytevar byte = 100
	fmt.Println(bytevar)

	//////////////////// Numeros reais
	var floatvar float32 = 10000000000000000000000000000000000000.12312123123
	fmt.Println(floatvar)

	var floatvar2 float64 = 100000000000000000000000000000000000001312312312311231231231231231231232434534343434.12312123123
	fmt.Println(floatvar2)

	//////////////////// Tipo string
	var stringvar string = "Texto"
	fmt.Println(stringvar)

	//////////////////// Tipo Char
	charvar := 'A'
	fmt.Println(charvar)

	//////////////////// Tipo Boolean
	var boolvar bool = true
	fmt.Println(boolvar)

	//////////////////// Tipo Error
	var errorvar error = nil
	var errorvar2 error = errors.New("Erro interno")
	fmt.Println(errorvar)
	fmt.Println(errorvar2)
}
