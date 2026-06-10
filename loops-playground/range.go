package main

import (
	"fmt"
)

// case 1: slicing using range:

func main() {
	fmt.Println("before ignoring index")

	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Println(index, value)
	}

	//case 2 : ignore index:
	fmt.Println("after ignoring index")
	for _, value := range nums {
		fmt.Println(value)
	}

	fmt.Println("ignoring value")
	//case 3: ignore value
	for index, _ := range nums {
		fmt.Println(index)
	}
	//case 4: range over string
	word := "Amaan"
	for index, char := range word {
		fmt.Println(index, char)
	}
	//above thing prints numbers because char is a rune
	//if you want to print the actual thing
	newWord := "ayaan"
	for _, char := range newWord {
		fmt.Println(string(char))
	}
	//case 5: maps: you can iterate over maps quick and easy
	ages := map[string]int{
		"Amaan": 20,
		"John":  25,
	}
	for key, value := range ages {
		fmt.Println(key, value)
	}

	//nested range: (TODO: )

	//case 7: modifying values:
	numbers := []int{1, 2, 3}
	for _, value := range numbers {
		value *= 10
		fmt.Println(value)
	}
}



//so index, value. and index, char. and
