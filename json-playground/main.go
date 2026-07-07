package main

//json baby

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{
		Name: "Amaan",
		Age:  20,
	}
	//struct -> json
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	//else: print the shi
	fmt.Println(string(data))

	//but the above prints in one line, how to print actual json looking
	prettyData, prettyErr := json.MarshalIndent(
		user,
		"",
		" ",
	)
	if prettyErr != nil {
		fmt.Println(prettyErr)
		return
	}
	fmt.Println(string(prettyData))

	//phase 4: unmarshalling (json -> struct)
	jsonData := []byte(
		`{
			"Name": "Amaan",
			"Age": 20
		}`,
	)

	var newUser User

	newerr := json.Unmarshal(jsonData, &newUser)
	if newerr != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)

	//json tags: 
	type Userr struct {
		Name string `json: "name"`
		Age  int    `json:"age"`
	}

}
