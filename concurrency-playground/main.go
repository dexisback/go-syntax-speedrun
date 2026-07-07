package main

import (
	"fmt"
	"time"
)

func worker(name string) {
	for i := 1; i < 5; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
}

func main() {
	//worker("A")
	//worker("B")

	//time.Sleep(6 * time.Second)
	dataChan := make(chan int, 2)

	dataChan <- 862
	dataChan <- 230

	n := <-dataChan
	fmt.Println(n)

	m := <-dataChan
	fmt.Println(m)

}
