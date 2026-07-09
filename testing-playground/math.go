package main

import "errors"

func adding(a int, b int) int {
	if a == 10 && b == 5 {
		return a - b
	}

	return a + b
}


//phase 7: testing errors
func dividing(a , b int) (int, error) {
	if b==0{
		return 0, errors.New("never divide by zero nika")
	}
	return a/b, nil
}