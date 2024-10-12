package main

import "fmt"

// Struct users
type user struct {
	name    string
	age     int
	address address
}

// Struct address of users
type address struct {
	street string
	number int
	zip    int
}

func main() {

	userAddress := address{street: "Rua colatina", number: 283, zip: 29102841}

	var user1 user
	user1.name = "Debora"
	user1.age = 26
	user1.address = userAddress

	user2 := user{name: "Adrian", age: 23, address: userAddress}

	fmt.Printf("Users: %v, %v\n", user1, user2)
}
