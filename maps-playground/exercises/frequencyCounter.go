package main

import (
	"fmt"
)

func countFrequency() {
	input := "mississipi"
	//output should be m->1, i-> 5, s-> 4, p-> 2
	//create a map where character (key)=> frequency (value)
	//iterate over every char of the strin
	//for each character, count[char]++
	counts := make(map[rune]int)
	//iterate over the array itself
	for _, char := range input {
		counts[char]++
	}

	//print the map:
	for char, count := range counts {
		fmt.Println(string(char), "->", count)
	}
	//string(char) because char is a rune and so if we print it alone then it would print ascii value, not the letter in itself

}
