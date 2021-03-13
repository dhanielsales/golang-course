package main

import (
	"fmt"
)

type user struct {
	name string
	age  uint
}

func (user user) save() {
	fmt.Println(user)
	fmt.Printf("Salvando os dados no banco de dados => %s \n", user.name)
}

func (user *user) update(newUserData user) {
	user.name = newUserData.name
	user.age = newUserData.age

	fmt.Println("Dados atualizados")
}

func main() {
	user1 := user{name: "Dhaniel", age: 26}
	user2 := user{name: "Felipe", age: 26}

	user1.save()
	user2.save()

	fmt.Println("---------------------------")

	user1.update(user{name: "Dhaniel Sales", age: user1.age})
	fmt.Println(user1)
}
