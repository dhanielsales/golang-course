package main

import "fmt"

type people struct {
	firstName string
	lastName  string
	age       int16
}

type student struct {
	people
	grade int8
}

func main() {
	student1 := student{people: people{"Felipe", "Castro", 20}, grade: 20}

	fmt.Println(student1)
	fmt.Println(student1.firstName)
	fmt.Println(student1.people.firstName)
	fmt.Println(student1.grade)
}
