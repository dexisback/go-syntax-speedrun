package main

import (
	"errors"
)

func widthdrawBal(b, a int) (int, error) {
	//takes in a balance and an amount as input , and returns integer and error
	//newBal = balance - amount
	if b < a {
		return b, errors.New("balance is lesser than the amount asked to widthdraw")
	}
	//else
	nB := b - a
	return nB, nil

}
