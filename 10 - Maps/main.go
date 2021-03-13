package main

import "fmt"

func main() {
	user := map[string]string{
		"name": "Dhaniel",
		"age":  "16",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"first": "Dhaniel",
		},
	}

	fmt.Println(user2["name"]["first"])

	delete(user, "age")
	delete(user2["name"], "first")

	fmt.Println(user)
	fmt.Println(user2)

}
