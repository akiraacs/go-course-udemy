package main

import "fmt"


func signalInverterWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 10
	fmt.Println(number)

	signalInverterWithPointer(&number)
	fmt.Println(number)
}