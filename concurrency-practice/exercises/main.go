package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(serviceName string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(serviceName, "starting...")
	time.Sleep(2 * time.Second)
	fmt.Println(serviceName, "ready✅")
}

//we'll call redis , etc etc

// func main() {
// 	//declare:
// 	var wg sync.WaitGroup
// 	services := []string{
// 		"Database",
// 		"Redis",
// 		"API",
// 		"File system",
// 	}

// 	wg.Add(len(services))

// 	for _, services := range services {
// 		go worker(services, &wg)

// 	}
// 	wg.Wait()
// 	fmt.Println("all services ready")

// }

// func squarePrinter(number int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("the square of ", number, "is ", number*number)

// }

// func main() {
// 	var wg sync.WaitGroup
// 	localDb := []int{
// 		1, 2, 3, 4, 5,
// 	}

// 	wg.Add(len(localDb))
// 	for _, r := range localDb {
// 		//form goroutines and execute them and outside of this loop execute wg.Wait() so that it waits until the complete goroutines are done
// 		go squarePrinter(r, &wg)

// 	}
// 	wg.Wait()

// 	fmt.Println("done")
// }

//instead of 1000 increments, make 10 goroutines and each increments 1000 times.
//so 10k total . first run without mutex and then run with mutex

//without mutex: this would print mostly 10k but sometimes randomly 7.6k , or something random. because multiple goroutines are trying to alter the same counter variable at the same time
// var counter int

// func incrementer(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 1000; i++ {
// 		counter++
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go incrementer(&wg)
// 	}

// 	wg.Wait()
// 	fmt.Println("final counter", counter)
// }

//with mutexes (correct):

// var (
// 	counter int
// 	mu      sync.Mutex
// )

// func increment(wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := 0; i < 1000; i++ {
// 		mu.Lock()
// 		counter++
// 		mu.Unlock()

// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go increment(&wg)
// 	}

// 	wg.Wait()

// 	fmt.Println(counter)
// }
// //now it prints a constant 10k everytime

//problem 2 -- now i gotta make hundreds of transactions run through multiple goroutines for a dummy bank balance and write code somehow that the balance gets protected -- one goroutine at a time

// var (
// 	balance = 1000
// 	mu      sync.Mutex
// 	wg      sync.WaitGroup
// )

// func Desposit(amount int) {
// 	defer wg.Done()
// 	mu.Lock()
// 	defer mu.Unlock()

// 	balance += amount
// 	fmt.Println("deposited:", amount, "Balance:", balance)

// }

// func withDraw(amount int) {
// 	defer wg.Done()
// 	mu.Lock()
// 	defer mu.Unlock()
// 	if balance >= amount {
// 		balance -= amount
// 		fmt.Println("withdrawn", amount, "balance remains:", balance)
// 	} else {
// 		fmt.Println("cannot withdraw as the current balance is less than withdraw amount ")
// 	}

// }

// func main() {
// 	wg.Add(6)
// 	go Desposit(1000)
// 	go withDraw(20)
// 	go Desposit(400)
// 	go withDraw(120)

// 	go Desposit(2308)
// 	go withDraw(293)

// 	wg.Wait()

// 	fmt.Println("final balance is :", balance)
// }

//problem 3 -- inventory : initial stock is 50, 20 customers buy the same item. so the final stock MUST be 30, regardless of races

var (
	stock = 50
	mu    sync.Mutex
	wg    sync.WaitGroup
)

func buyItem(customerID int) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()

	if stock > 0 {
		stock--
		fmt.Println("customer %d bought an item, stock left is : %d", customerID, stock)
	} else {
		fmt.Println("customer %d is out of stock!", customerID)
	}

}

func main() {
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go buyItem(i)
	}
	wg.Wait()
	fmt.Println("final stock is:", stock)
}
