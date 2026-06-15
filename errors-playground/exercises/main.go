package main

import (
	"fmt"
)

func main() {
	// result, err := divide(10, 0)
	// if err != nil {
	// 	wrappedError := fmt.Errorf("division failed: %w", err)
	// 	fmt.Println(wrappedError)
	// 	return
	// }

	// //else:
	// fmt.Println(result)

	// result, err := findUser(1)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// //else print error:
	// fmt.Println(result)

	// result, err := validateAge(17)
	// // resultOne, errOne := validateAge(20)

	// if err != nil {
	// 	//mtlb error hai code me
	// 	fmt.Println(err)
	// }
	// //else print result
	// fmt.Println(result)
	result, err := widthdrawBal(1000, 500)
	if err != nil {
		fmt.Println(err)

	}

	//else
	fmt.Println(result)
}
