//phase 1: sequentially executing stuff:

package main

import (
	"fmt"
	"sync"
)

// func task(name string) {
// 	for i := 1; i <= 3; i++ {
// 		fmt.Println(name, i)
// 		time.Sleep(time.Second) //learn
// 	}

// }

// func main() {
// 	task("A")
// 	task("B")
// }

// // goroutine:
// func main() {
// 	go task("A")
// 	time.Sleep(4 * time.Second)
// }

func task(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("working...")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go task(&wg)
	wg.Wait()

}
