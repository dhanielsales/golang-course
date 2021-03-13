package main

import "fmt"

type user struct {
	name     string
	age      uint8
	address  address
	address2 address
	address3 address
}

type address struct {
	street string
	number uint8
}

func main() {
	var myUser user

	myUser.name = "Dhaniel"
	myUser.age = 20
	fmt.Println(myUser)

	// myUser2 := user{"Dhaniel Ricardo", 11, address{"Rua A", 12}}
	// fmt.Println(myUser2)

	myUser3 := user{name: "Dhaniel Ricardo", age: 11, address: address{street: "Rua B", number: 123}}
	fmt.Println(myUser3)
}
