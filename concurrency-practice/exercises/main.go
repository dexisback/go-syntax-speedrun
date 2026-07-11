package main

import (
	"fmt"
	"sync"
)

// func worker(serviceName string, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	fmt.Println(serviceName, "starting...")
// 	time.Sleep(2 * time.Second)
// 	fmt.Println(serviceName, "ready✅")
// }

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

// var (
// 	stock = 50
// 	mu    sync.Mutex
// 	wg    sync.WaitGroup
// )

// func buyItem(customerID int) {
// 	defer wg.Done()

// 	mu.Lock()
// 	defer mu.Unlock()

// 	if stock > 0 {
// 		stock--
// 		fmt.Println("customer %d bought an item, stock left is : %d", customerID, stock)
// 	} else {
// 		fmt.Println("customer %d is out of stock!", customerID)
// 	}

// }

// func main() {
// 	for i := 1; i <= 20; i++ {
// 		wg.Add(1)
// 		go buyItem(i)
// 	}
// 	wg.Wait()
// 	fmt.Println("final stock is:", stock)
// }

// ----rw---- //

//scores need to be a global variable:

// var (
// 	scores = map[string]int{
// 		"Amaan": 100,
// 	}

// 	rw sync.RWMutex
// 	wg sync.WaitGroup
// )

// func reader(id int) {
// 	defer wg.Done()

// 	for i := 0; i < 5; i++ {
// 		rw.RLock()
// 		fmt.Printf("Reader %d -> %d\n", id, scores["Amaan"])
// 		rw.RUnlock()

// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

// func writer() {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(time.Second)

// 		rw.Lock()
// 		scores["Amaan"]++
// 		fmt.Println("Writer updated:", scores["Amaan"])
// 		rw.Unlock()
// 	}
// }

// func main() {
// 	go writer()

// 	for i := 1; i <= 10; i++ {
// 		wg.Add(1)
// 		go reader(i)
// 	}

// 	wg.Wait()
// }

//problem 3: building a small cache :

//own tiny in memory redis:

//get() for reading , set() for writing , RWMutex so many readers can read together, but writes are exclusive

// var (
// 	cache = make(map[int]string)
// 	rw    sync.RWMutex
// )

// func GET(key int) string {
// 	rw.RLock()
// 	defer rw.RUnlock()
// 	//function
// 	return cache[key]

// }

// func SET(key int, value string) { //returns nothin
// 	rw.Lock()
// 	defer rw.Unlock()
// 	//function
// 	cache[key] = value
// }

// func main() {
// 	SET(1, "Amaan")
// 	SET(2, "Ayaan")
// 	fmt.Println(GET(1))
// 	fmt.Println(GET(2))
// }

//problem 4 -- building an in memory cache in a backend

// type User struct {
// 	Name string
// }

// var (
// 	cache = make(map[int]User)
// 	rw    sync.RWMutex
// )

// func GETUSER(id int) (User, bool) {
// 	rw.RLock()
// 	defer rw.RUnlock()

// 	user, ok := cache[id]
// 	return user, ok
// }

// func UPDATEUSER(id int, user User) {
// 	rw.Lock()
// 	defer rw.Unlock()

// 	cache[id] = user

// }

// func main() {
// 	UPDATEUSER(1, User{Name: "amaan"})
// 	UPDATEUSER(1, User{Name: "ayaan"})

// 	if user, ok := GETUSER(1); ok {
// 		fmt.Println(user.Name)
// 	}
// 	UPDATEUSER(1, User{Name: "Amaan Siddiqui"})

// 	if user, ok := GETUSER(1); ok {
// 		fmt.Println(user.Name)
// 	}
// }

//worker pool exercise for channels:

//classic worker pool pattern:
// 1 producer (some api service ) creates jobs
// 1 shared jobs channel
// 3 workers compete for jobs
// every job is processed exactly once

// func worker(id int, jobs <- chan int, wg *sync.WaitGroup ){
// 	defer wg.Done()

// 	for job := range jobs {
// 		fmt.Println("worker", id, "processing job", job)
// 		time.Sleep(2 * time.Second) // simulate some work here

// 	}
// }

// func main(){
// 	var wg sync.WaitGroup

// 	jobs := make(chan int) //create a jobs channel

// 	//start 3 workers
// 	for i :=1; i<=3 ; i++{
// 		//sparn the function 3 times (with goroutines ), and obviously add each time 1 goroutine watchgroup
// 		wg.Add(1)
// 		go worker(i, jobs, &wg)

// 		//producer: creates 20 jobs:
// 		for i:=1 ; i<=20 ; i++ {
// 			jobs <- i
// 		}
// 		close(jobs)

// 		wg.Wait()

// 	}
// }

//building a pipeline/ simulating a pipeline:
//generate -> square -> print

func generate(ch1 chan<- int) {
	defer close(ch1)

	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
}

func square(ch1 <-chan int, ch2 chan<- int) {
	defer close(ch2)

	for num := range ch1 {
		ch2 <- num * num
	}
}

func printResult(ch2 <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for result := range ch2 {
		fmt.Println(result)
	}
}

func main() {
	var wg sync.WaitGroup

	ch1 := make(chan int)
	ch2 := make(chan int)

	go generate(ch1)
	go square(ch1, ch2)

	wg.Add(1)
	go printResult(ch2, &wg)

	wg.Wait()

}
