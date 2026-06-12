package main

//prints only the unique characters from inputted things

import (
	"fmt"
)

func uniquePrinter(s []int) {
	//takes in an input of an array
	//goes through the array
	//same concept -> map stores only till count:1, and then we print out the map
	counts := make(map[int]int)

	for _, value := range s {
		_, exists := counts[value]
		if exists {
			continue //skip that value
		}
		counts[value]++
	}
	//now we need to print out the map
	fmt.Println(counts) //this prints out map
	//better: we print out the values {arrays} , we need a loop again for dat
	for key := range counts {
		fmt.Println(key)
	}
}



//NOTE: THIS PRINTS OUT IN RANDOM VALUES, BECAUSE MAPS DONT STORE IN ORDER 
//IF WE WANT ORDER:
// seen := make(map[int]bool)

// for _, value := range s {
//     if !seen[value] {
//         fmt.Println(value)
//         seen[value] = true
//     }
// }