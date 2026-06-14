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
	newVariable := map[int]User{
		1: {
			Name: "amaan",
			Age:  21,
		},
		2: {
			Name: "Ayaan",
			Age:  13,
		},
	}
	_ = newVariable

	//nested structs
	type Address struct {
		City string
	}

	type User struct {
		Name    string
		Address Address
	}

	amaanUser := User{
		Name: "amaan",
		Address: Address{
			City: "bilaspur",
		},
	}

	fmt.Println(amaanUser)
	fmt.Println(amaanUser.Address.City)

	//anonymous structs -- useful for quick experiemtns:
	Anonymoususer := struct{
		Name string 
		Age int
	}{
		Name: "Amaan",
		Age: 21,
	}

	_ = Anonymoususer
	



}
