package main

import (
	"errors"
)

func validateAge(age int) (bool, error) {
	if age < 18 {
		return false, errors.New("age below 18!")
	}

	//else:
	return true, nil
}
