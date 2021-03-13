package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ARRAYS
	array1 := [5]int{}
	fmt.Println(array1)

	array2 := [5]string{}
	array2[2] = "Posição 3"
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3}
	fmt.Println(array3)

	// SLICES
	slice1 := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice1)

	// VERIFICAÇÂO DE TIPOS
	fmt.Println(reflect.TypeOf(array1), reflect.TypeOf(slice1))

	// SLICES
	slice1 = append(slice1, 16)
	fmt.Println(slice1)

	slice2 := array1[0:3]
	fmt.Println(slice2)

	slice3 := array1[0:]
	fmt.Println(slice3)

	slice3[2] = 12
	fmt.Println(slice3)

	// ARRAYS INTERNOS
	fmt.Println("-------------------")
	slice4 := make([]float32, 10, 11)

	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 11)
	slice4 = append(slice4, 12)

	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice5 := make([]float32, 10)
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
	slice5 = append(slice5, 11)
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
}
