package main

import (
	"fmt"
	"time"
)

// // phase 1 : goroutines
// //done

// // now ,

// func work(name string) {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(name, i)
// 		time.Sleep(time.Second) //time.Sleep pauses the current goroutine for a specific duration
// 		//time.Sleep(2 * time.Second ) means wait for 2 seconds and then go below the line
// 		//time.Sleep(time.Second ) means wait for 1 second

// 	}
// }
// func main() {
// 	go work("A")
// 	work("B")
// }




//phase 2: waitgroups:

//func main(){
//	go work("A")
//} me programs exits before the goroutine finishes 


//earlier we fixed it using time.Sleep() but proper proper we need waitgroups

//wait until all my goroutines finish

