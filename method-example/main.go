package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) save() {
	fmt.Println("Saved user!")
}

func (u *user) haveABirthday() {
	u.age++
}

func main() {
	user1 := user{name: "Adrian", age: 23}
	fmt.Println(user1)

	user1.save()

	user1.haveABirthday()
	fmt.Println(user1.age)
}