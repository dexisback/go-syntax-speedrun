package main

import "fmt"

func greeting() string {
	return "hello"
}

func main() {
	fmt.Println(greeting())
}


func Adder[T int64 | float64](a T , b T)T{
	return a + b 
}

