package main

import "sync"

//1: declaring

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1) //so we tell the thing ki bhai only one goroutine we r gonna spawn
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("hellow")
// 	}()

// 	wg.Wait()
// }

// 2: loops: you'll use these for basically anything iterable
// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			fmt.Println(i)
// 		}(i)

// 		wg.Wait()
// 	}
// }

// // 3: named worker function:
// func process(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println(id)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go process(i, &wg)

// 	}
// 	wg.Wait()
// }

//main creates ONE waitGroup, this is literally
//the ONLY WaitGroup in the entire program

//&wg means we pass the address of that waitGroup
//wg *sync.WaitGroup means "i expect the address of a WaitGroup"

//pattern 3: named worked function:
//This pattern is extremely common in Go because synchronization primitives (WaitGroup, Mutex, Cond, etc.) are shared objects, not things you copy.

// func process(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println(id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)

// 		go process(i, &wg)
// 	}

// 	wg.Wait()

// }

//pattern 4: useless

//pattern 5: different workers but the same &wg

// func main() {

// 	wg.Add(3)
// 	go fetchUsers(&wg)
// 	go fetchOrders(&wg)
// 	go fetchProducts(&wg)

// 	wg.Wait()

// }

//pattern 6: if we dont know beforehand how many jobs exist


// func main(){
// 	jobs:= []string{
// 		"a",
// 		"b",
// 		"c",
// 		"d",
// 	}


// 	wg.Add(len(jobs))

// 	for _, job := range jobs {
// 		// go process(job , &wg)
// 	}

// 	wg.Wait()
// }
//pattern 7: error safe worker :

// func worker(wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	if somethingFailed {
// 		return 
// 	}

// 	//else , continue with your work like u were 
// }
