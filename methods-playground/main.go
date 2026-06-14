package main

import "fmt"

// first method:
type User struct {
	Name string
	Age  int
	Email string
}

// 1. first method:
func (u User) Greet() {
	fmt.Println("hellow there, ", u.Name)
}

// 2.returns some value
func (u User) isAdult() bool {
	return u.Age >= 18
}

// multiple methods:
func (u User) DisplayName() string {
	return "username is:" + u.Name
}

// different types can have same methods:
type Product struct {
	Name  string
	Price int
}

func (u User) NamePrinter() {
	fmt.Println(u.Name)
}

func (p Product) NamePrinter() {
	fmt.Println(p.Name)
}

// updating things ke liye we need pointers:
func (u *User) MinorMaker() int {
	return (u.Age - 18) //☠️
}


//BE example:
func (u User) Validate() bool {
	return u.Email != ""
}


func main() {

	user := User{
		Name: "amaan",
		Age:  20,
		Email: "texamaan86@gmail.com",
	}
	user.Greet()

	fmt.Println(user.isAdult())
	fmt.Println(user.DisplayName())

	fmt.Println(user.MinorMaker())
}
