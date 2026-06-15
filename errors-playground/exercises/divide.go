// handle the error of when a/b me b == 0
package main

import (
	"errors"
)

func divide(a, b int) (float64, error) {
	//take sin a and b and divides them
	if b == 0 {
		return 0, errors.New("b is 0")
	}
	//else
	answer := float64(a) / float64(b)
	return answer, nil
}
