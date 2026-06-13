package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{
		Name: "amaan",
		Age:  20,
	}
	fmt.Println(user)
	//print specific values from the struct
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	//modify fields
	user.Age = 22
	fmt.Println(user)
	//multiple users with the same struct (so, same schema basically)
	user1 := User{
		Name: "ayaan",
		Age:  23,
	}

	fmt.Println(user, user1)

	//structs in slices:
	users := []User{
		{Name: "amaan", Age: 20},
		{Name: "john", Age: 53},
		{Name: "ayaan", Age: 13},
	}
	fmt.Println(users)

	//loop over users:
	for _, user := range users {
		fmt.Println(user.Name)
	}

	//structs in maps:
	users := map[int]User{
		1: {
			Name: "amaan",
			Age: 21
		},
		2: {
			Name: "Ayaan",
			Age: 13
		},
	}

}
