//mashup of everythang

package main

type User struct {
	Name string
	Age  int
}

func FindUser(users []User, name string) (User, bool) {
	//takes in a slice of type User, and a name in the input
	//and returns the User if found, if not found then returns a bool
	//suppose users is:
	// inputUser := []User{
	// 	{Name: "amaan", Age: 20},
	// 	{Name: "ayaan", Age: 13},
	// }

	// inputName := "ayaan"

	//now we iterate through the inputUser, and if we find the inputName, then we return true, if not found we return false

	for _, values := range users {
		if values.Name == name {
			return values, true
		}
	}
	return User{}, false
}
