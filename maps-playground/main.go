package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi")
	// ages := map[string]int{}
	ages := map[string]int{
		"ayaan": 21,
		"anees": 25,
	}
	ages["amaan"] = 21
	ages["john"] = 25
	fmt.Println(ages)
	delete(ages, "john") //delete the thing
	fmt.Println(ages)
	fmt.Println(ages["john"]) //prints 0 because go returns the zero value of the value type
	age, exists := ages["aayesha"]
	fmt.Println(age)
	fmt.Println(exists)
	age, thisShouldReturnTrue := ages["ayaan"]
	fmt.Println(thisShouldReturnTrue)
	//later you will write something like
	// user , exists := users[userId]
	//if !exists {
	//	return 
	//}
	//nested maps:
	nestedMaps := map[string]map[string]string{
		"amaan" : {
			"email": "texamaan86@gmail.com",
		},
	} 
	fmt.Println(nestedMaps)

	//usually structs are preffered in prod

	//making an empty map:
	emptyMapUsingMake := make(map[string]int)
	_ = emptyMapUsingMake

	//this is equivalent to {}
	

}
