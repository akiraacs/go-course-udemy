package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

    // Declarando array e atribuindo valores depois
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

    // Declarando e atribuindo valores a um array
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

    // Outra forma de declarar e atribuir valores a um array só que sem informar a quantidade de valores
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

    // Slice é uma fatia de um array, ou melhor dizendo um ponteiro de um array
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)
    
	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	// Arrays Internos
	fmt.Println("----------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}