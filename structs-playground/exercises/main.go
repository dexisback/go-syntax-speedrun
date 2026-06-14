package main

import "fmt"

func main() {
	fmt.Println("hi")
	threeBooks()
	fmt.Println(findTopStudents())
	usingMaps()
	users := []User{
		{Name: "amaan", Age: 20},
		{Name: "ayaan", Age: 13},
		{Name: "anees", Age: 50},
	}
	user, found := FindUser(users, "ayaan")

	if found {
		fmt.Println("user found ->", user)
	} else {
		fmt.Println("user not found sorrry")
	}
}
