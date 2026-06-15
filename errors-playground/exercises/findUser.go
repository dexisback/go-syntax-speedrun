// if id == 1 then only exists and return true else return error of user not found
package main

import (
	"errors"
)

func findUser(id int) (int, error) {
	if id == 1 {
		return id, nil

	}

	//else error wala route
	return 0, errors.New("id is not in the db")
}
