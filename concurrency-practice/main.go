package main

import (
	"fmt"
	"sync"
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

// func work(wg *sync.WaitGroup) {
// 	defer wg.Done() //equivalent to writing wg.Done() and then return

// 	for i := 1; i <= 5; i++ {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup //creates a waitgroup (a counter of unifinished goroutines)
// 	wg.Add(1) //one goroutine is about to start we tell the language

// 	go work(&wg) //pass the workgroup by pointer. why pointer? because both main and work() must modify the same workgroup

// 	wg.Wait()  //dont continue until counter becomes 0
// 	fmt.Println("finished ✅")
// }

// //so initially "counter = 0". then we do wg.Add(1) means counter became 1. after wg.Done() counter decreased and became 0 (and wg.Wait is what holds the compiler to execute uske neecehe wali lines)

// func worker(name string, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	fmt.Println(name, "started")
// 	time.Sleep(3 * time.Second) //simulating some work over here
// 	fmt.Println(name, "finished ")

// }

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go worker("A", &wg)
// 	go worker("B", &wg)

// 	wg.Wait()

// 	fmt.Println("main exits...")
// }
//to observe is -- both workers run together -> main waits for BOTH -> only then main does exit

//common mistakes -- if you forget wg.Done() , then the program waits forever
//

//phase 3 -- race conditions + mutex
//everything up until now was goroutines, now we will make the goroutine touch the same variable

func increment(counter *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mu.Lock()
		(*counter)++
		mu.Unlock()
	}
}

func main() {
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)

	go increment(&counter, &mu, &wg)
	go increment(&counter, &mu, &wg)

	wg.Wait()

	fmt.Println(counter)
}


//usually you need mu.Lock()
//cache[id] = user
//mu.Unlock()

//to prevent race condition


//defer Unlock()
//instead of mu.Lock() and mu.Unlock() and typing the code in b/w that area
//you can simply write defer mu.Unlock() and go will understand what you mean

