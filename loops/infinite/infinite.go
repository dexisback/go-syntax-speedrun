package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
}