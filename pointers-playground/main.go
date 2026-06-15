package main

import "fmt"

func main() {
	i, j := 42, 501
	fmt.Println(i, j)   //prints out the actual values
	fmt.Println(&i, &j) //print out the address of em

	//now lets say i store the address of variable 'i' in something , say p
	p := &i
	fmt.Println(p) //this simply prints out the address of i
	//deferencing:
	fmt.Println(*p) //this cancels out the address & effect,and so we will get the value again

	//modfying something through pointers
	age := 21
	fmt.Println("before", age)
	ptr := &age

	*ptr = 50
	fmt.Println("after pointer", age)
	//so youre not changing it directly

	//methods needed pointers whenever we udpating something

}

type User struct {
	Name string
	Age  int
}

// passing struct by pointers:
func UpdateUser(user *User) {
	user.Name = "Amaan"
}

// methods with pointer receivers:
func (u *User) Rename(name string) {
	u.Name = name
}

//nill pointers:
// var userButEmpty *User
// if user == nill {
// 	//handle logic
// }

// optional fields in APIs:
type UpdateUserRequest struct {
	Name *string `json: "name"`
	Age  *int    `json: "age"`
}

// returning pointers
func FindUser(id int) (*User, error)

// why pointers avoid copying :
func ProcessUser(u User)  //copies the whole struct
func ProcessUser(u *User) //passes an address , this matters for large files

// note: slices + maps are already reference like
func AddUser(users *[]User) //you need not add the pointer because it is pre built for slices and maps in go
func AddUser(users *[]User) //is enough

func (s *UserService) GetUser(id int) (*User, error) {
	return s.repo.FindByID(id)
}
