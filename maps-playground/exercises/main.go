package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi")
	fmt.Println(duplicateDetect([]int{1, 1, 23}))
	uniquePrinter([]int{1, 1, 2, 3, 4, 4})
	answer := (firstNonRepeatingVal("amaan"))
	fmt.Println(string(answer))
	
}
