package main

import (
	"fmt"
)

// "sync"
// "time"
// "sync"
// "time"

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

// func increment(counter *int, mu *sync.Mutex, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := 0; i < 1000; i++ {
// 		mu.Lock()
// 		(*counter)++
// 		mu.Unlock()
// 	}
// }

// func main() {
// 	counter := 0
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex

// 	wg.Add(2)

// 	go increment(&counter, &mu, &wg)
// 	go increment(&counter, &mu, &wg)

// 	wg.Wait()

// 	fmt.Println(counter)
// }

//usually you need mu.Lock()
//cache[id] = user
//mu.Unlock()

//to prevent race condition

//defer Unlock()
//instead of mu.Lock() and mu.Unlock() and typing the code in b/w that area
//you can simply write defer mu.Unlock() and go will understand what you mean

//RW Mutexes:
//there will be no blockage to reading but only writing

// var (
// 	cache = map[int]string{
// 		1: "Amaan",
// 	}
// 	rw sync.RWMutex
// )

// func read(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	rw.RLock()
// 	fmt.Println(cache[id])
// 	rw.RUnlock()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(3)

// 	go read(1, &wg)
// 	go read(1, &wg)
// 	go read(1, &wg)
// }
// //all three readers can run together with this

// //writing agar karna ho with rw locks then?:

// func write(id int, name string , wg *sync.WaitGroup){
// 	defer wg.Done()

// 	rw.Lock()

// 	cache[id] = name

// 	rw.Unlock()
// }

// //so we use normal rw.Lock() and rw.Unlock() when its about a write, but RLock and RUnlock when its about a read

//mixed reads and writes:

// var (
// 	counter int
// 	rw      sync.RWMutex
// )

// func reader(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	rw.RLock() //normal
// 	fmt.Println("reader with id:", id, "reads", counter)

// 	time.Sleep(time.Second)

// 	rw.RUnlock()

// }

// func writer(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	counter++

// 	fmt.Println("write updated man :)")
// 	time.Sleep(time.Second)

// 	rw.Unlock()
// }

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(4)

// 	go reader(1, &wg)
// 	go reader(2, &wg)
// 	go reader(3, &wg)

// 	//write:
// 	go writer(&wg)

// 	wg.Wait()

// }

//channels:

//creating a channle -- done
// send and receive -- done
// why goroutine? -- done
//blocking nature -- done
//multiple sends -- done
//buffered channels -- done
//closing channel -- done
//comma , ok : you add an "ok" after passing a channel.
//that "ok" will give you true or false based on if the channel is closed or not

//if channel is closed -> ok ===false

// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		// ch <- 80
// 		close(ch)
// 	}()
// 	value, ok := <-ch

// 	if ok == false {
// 		fmt.Println("channel closed")
// 	}
// 	//else:
// 	fmt.Println(value)
// }

//backend example: read image -> resize image -> upload image

//----------------------------------

//fan out: one producer multiple workers. workers compete for jobs

//worker takes an id, a job which is a channel and waitgroup so that when we call it we can pass the addres sby "&"

// func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for job := range jobs {
// 		fmt.Printf("worker %d processing job %d\n", id, job)
// 		time.Sleep(time.Second) // sleep for one second
// 		fmt.Printf("worker %d finished job %d\n", id, job)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	jobs := make(chan int)

// 	//spawn 3 workers

// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1)
// 		go worker(i, jobs, &wg)
// 	}
// 	//producer:
// 	for i := 1; i <= 3; i++ {
// 		jobs <- i
// 	}

// 	close(jobs)

// 	wg.Wait() //wait until the above block is done
// }

//fan in: multiple workers. one results channel

//todo:

//select:
//todo

//select is basically go's way of saying , i have multiple channels, whichever one becomes ready first, we do that

//exercise 1: fastest gorotuine wins:

// func main() {
// 	channel1 := make(chan string)
// 	channel2 := make(chan string)

// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		channel1 <- "Database finished"
// 	}()

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		channel2 <- "redis finished"
// 	}()

// 	//select ka use case yaha pe:
// 	select {
// 	case msg := <-channel1:
// 		fmt.Println(msg)

// 	case msg := <-channel2:
// 		fmt.Println(msg)
// 	}
// }

// //we wait for whichever event happens first
// //if you're buildign a backend. two things can happen -> db returns data or db returns an error
// //and we dont knwo which happens first
// //so we do userChan and erroChan
// //instead of simply doing user := <- userChan //because waits here forever if there is no response and there instead is error

//exercise 2:custom timeout:
//make channel1 sleep and channel2 sleep 3 seconds

// func main() {
// 	ch := make(chan string)

// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		ch <- "done"
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println(msg)

// 	case <-time.After(2 * time.Second):
// 		fmt.Println("timed out")
// 	}
// }

// //what we r doing is giving the db 5 seconds (dummy work, suppose db takes up 5 seconds to boot up)
// //and we r saying ki bhai if something exceeds beyond 2 seconds then we will print out "timed out" , even if the process is going on behind we dont care
// //so that we built

//exercise 4: two producers:

// func producer(name string, delay time.Duration, ch chan string) {
// 	time.Sleep(delay)
// 	ch <- name
// }

// func main() {
// 	channel := make(chan string)
// 	go producer("prod 1 ", 2*time.Second, channel)
// 	go producer("prod 2 ", 2*time.Second, channel)

// 	fmt.Println(<-channel)
// }

//exercise 5:

// func main() {
// 	userChan := make(chan string)
// 	errorChan := make(chan string)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		userChan <- "Amaan"
// 	}()

// 	go func() {
// 		time.Sleep(time.Second)
// 		errorChan <- "db failed"
// 	}()

// 	select {
// 	case user := <-userChan:
// 		fmt.Println(user)

// 	case err := <-errorChan:
// 		fmt.Println(err)

// 	}

// }

//we get db failed (error) because we designed in such a way that the error arrives first, so this is backend technique ki bhai if erorr comes up first then receive the thing and print errors out. so this is go's way of printing errors first

//default:

// func main() {
// 	channel := make(chan string)

// 	select {
// 	case v := <-channel:
// 		fmt.Println(v)

// 	default:
// 		fmt.Println("nothin available")

// 	}

// }
