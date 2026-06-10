package main

import (
	"fmt"
	
)

func stringReverser(s string) {

	//takes a string, iterates completely on it and then
	//returns a string
	//looping backwards : option a
	//akshun appending in a new array and printing that out: option b
	//option a below:

	for i := len(s) - 1; i >= 0; i-- {
		fmt.Printf(string(s[i]))
	}

}
