package main

import "fmt"

// import ("fmt")

// func multiplier(n int) string {
// 	//strings of numbers
// 		for i:= 1; i<=10 ; i++ {
// 			fmt.Printf()
// 		}

// }

func table(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}
