package main

import (
	"errors" //wouuuuw lego amazed meme
	"fmt"
)

// //errors are values, not exceptions
// //so we just return them like any other value

// //pehle we did:
// func getUser() (string, int) {
// 	return "Amaan", 20
// }

// //now we extend that to
//
//	func getUserr() (string, error){
//		//code logic
//	}

func getUser(id int) (string, error) {
	if id != 1 {
		return "", errors.New("user not found")
	}

	return "Amaan", nil
} //success -> returns username + nil , failure -> returns empty string + actual errro

func main() {
	//first error

	user, err := getUser(2)
	fmt.Println(user)
	// fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		// return  // need to remove this return because this will completely exit out of main and the bottom lins never reach bottom loc, so its a fundamental choice you make whether the error is so much that you want your app to crash or continue but still print out the error
	}
	//success paths:
	CorrectUser, err := getUser(1)

	if err != nil {
		fmt.Println(err)
		return
	}
	//else
	fmt.Println(CorrectUser)
	//phase 7 : custom errors:done (basically errors.New("your error"))

	//phase 9: wrapping errors
	fmt.Errorf("failed to fetch errors", err)
	fmt.Errorf("database query failed ", err)

	//phase 10: 
}

//phase 8: propagating errors:

func getUserrr() (string, error) {
	//imagine this talking to postgres, something fails and it returns "", err
	return "", errors.New("db failed")
}

func handleUser() error {
	_, err := getUserrr()

	if err != nil { // because getUserrr() will return an err if something goes wrong
		return err
	}
	//else
	return nil
}


type User struct {
	Name string 
	Age int 
	Bal int
}
func userFinder(id int) (User, error){
	//finds the user from the given id in db
}


func mainn(){
	user, err := userFinder(id)
	if err != nil {
		return err
	}
}

//go preferes returning early than aage jaane dena;
func CreateUser(email string) error {
	if email = "" {
		return errors.New("email required!!")
	}

	if len(email) < 5 {
		return errors.New("invalid email")
	}

	return nil
}