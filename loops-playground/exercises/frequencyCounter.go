package main

import "fmt"

func countFrequency(s string) {
	// given a string, count the frequency of all the letters in it
	// maps are go's version of {a: 3, n: 2, b: 1}
	counts := make(map[rune]int)

	for _, char := range s {
		counts[char]++
	}

	for char, count := range counts {
		fmt.Println(string(char), "->", count)
	}
}
