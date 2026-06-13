package main

func firstNonRepeatingVal(s string) rune {
	//takes in a string s and iterates through it,
	//go through each character and mpa its frequency
	//find the first character whose frequency is 1
	//since maps are unordered you will have to go through the string again and ask the map, ki iski frequency 1 hai kya
	trackMap := make(map[rune]int)

	//pass 1: fill the map with char and their frequencies
	for _, char := range s {
		trackMap[char]++
	}

	//pass 2: ifnd first character w frequency of 1
	for _, char := range s {
		if trackMap[char] == 1 {
			return char
		}
	}

	return 0

}
