package main

import "fmt"


func main() {

    var1, var2 := 1, 2

    // Formas de criar variavel pointeiro
    // pointerVar1 := &var1

    var pointerVar1 *int
    pointerVar1 = &var1

    fmt.Printf("Memory address var1: %v\n", pointerVar1)

    // Atribui valor para var1 atraves de seu ponteiro
    *pointerVar1 = 21

	fmt.Println(var1, var2)
	fmt.Println(*pointerVar1)// desreferenciação
}
