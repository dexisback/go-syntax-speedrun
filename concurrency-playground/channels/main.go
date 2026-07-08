// // Print numbers from another goroutine.
// // Sum numbers using workers.
// // Build a worker pool.
// // Build a concurrent web scraper.

// package main

// import ( "fmt"
// 	"sync"
// )

// //phase 1: channels ke saare apis
// //like wg channels ke bhi zyaada kuch methods zyaada nahi hote:

// //making a channel
// chForInt  := make(chan int) //create a portal that transports integers
// chForUserType := make(chan User)
// chForString := make(chan string)

// //sending :

// //recieving me var laga do saamne

package main

import (
	"fmt"
)

// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		ch <- 42 //sending via one goroutine
// 	}()

// 	num := <-ch
// 	fmt.Println(num)

// }

// #1: anonymous goroutines sending :
// func main() {
// 	ch := make(chan string)

// 	go func() {
// 		ch <- "hi"
// 	}()

// 	newVar := <-ch
// 	fmt.Print(newVar)
// }

// // #2: named worker
// func worker(ch chan int) {
// 	ch <- 100
// }

// func main() {
// 	ch := make(chan int)

// 	go worker(ch)

// 	fmt.Println(<-ch)
// }

// // #3: send and receive multiple vals:
// func main() {

// 	ch := make(chan int)
// 	go func() {
// 		ch <- 1
// 		ch <- 2
// 		ch <- 3
// 		ch <- 69
// 	}()

// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// // #4: loop se multiple vals receive krliye
// func main() {
// 	ch := make(chan int)
// 	go func() {
// 		ch <- 1
// 		ch <- 2
// 		ch <- 3
// 	}()

// 	for v := range ch {
// 		//receive in loop
// 		fmt.Println(v)
// 	}

// 	close(ch)

// }

// #5: passing structs in channels:
// type User struct {
// 	username string
// }

// func main() {
// 	chWithCustom := make(chan User)
// 	chWithCustom <- User{username: "Amaan"}

// 	receivingThingy := <-chWithCustom
// 	fmt.Println(receivingThingy)
// }

//#6: solving deadlocks: whenever deadlock happens means ki aapka same goroutine me sending and receiving hora

// func main() {
// 	dataChan := make(chan int)

// 	go func(){
// 		fmt.Println(<-dataChan)
// 	}()

// 	dataChan <- 862

// 	//new timeline:
// 	//create channel
// 	//create goroutine
// 	//receiver goroutine: waiting on <-ch
// 	//in main: we send 862

// }

// // #7: both sender and receiver in diff goroutines:
// func main() {
// 	newChannel := make(chan int)
// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	//sending ka goroutine:
// 	go func() {
// 		defer wg.Done()
// 		newChannel <- 900
// 	}()

// 	//receiving ka goroutine:
// 	go func() {
// 		defer wg.Done()
// 		receiving := <-newChannel
// 		fmt.Println(receiving)
// 	}()

// 	wg.Wait()
// }

// func main(){
// 	ch := make(chan int)

// 	go func(){
// 		ch <-900
// 	}()
// 	//after this line there are two goroutines (one main and one this sender)

// 	fmt.Println(<-ch) //now at this line main executes this line, and finds ki bhai abhi toh <- ch exist hi nahi krta
// 	//so main cannot exit right now. it says "i will wait here until someone sends me a value"

// }

// // buffered and unbuffered channel:
// func main() {
// 	dataChannel := make(chan int)

// 	dataChannel <- 730

// 	receivingThingy := <-dataChannel
// 	fmt.Println(receivingThingy)

// }

//above wont run because ek hi goroutine me sab hojara
//BUT theres something called buffered channel, in which basically for a temperory amount of time you can also assign SPACE
//inside of a channel. so they can convey the message bin the same goroutine aswell

// func main() {
// 	dataChannel := make(chan int, 2)

// 	dataChannel <- 730
// 	dataChannel <- 730

// 	receivingThingy := <-dataChannel
// 	receivingThingyOne := <-dataChannel
// 	fmt.Println(receivingThingy)
// 	fmt.Println(receivingThingyOne)

// }




// func main(){
// 	datachan := make(chan int)

// 	go func(){
// 		for i:= 0; i<1000; i++{
// 					datachan <- i

// 			}
// 				close(datachan)
// 	}
// 	()

// 	for n:= range datachan{
// 		fmt.Println(n)
// 	}

// }


// func main(){
// 	errChannel := make(chan error)
// 	//do some work that returns an error variable
// 	//assign the error to the variable 
// 	//and then pass it through the channel 
// 	errChannel <- err

// 	//receive it later in some place else 


// }