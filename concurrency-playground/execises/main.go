//phase 1: sequentially executing stuff:

package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second) //learn
	}

}

func main() {
	task("A")
	task("B")
}
