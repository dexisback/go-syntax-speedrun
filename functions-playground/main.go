package main

import "fmt"

// 4: shared type, done but
func add(a, b int) int {
	return a + b
}

//5:multiple return types:
func getUserInfo() (string, int) {
	return "Amaan", 20
}

// 7: basically all backend code:
// func findUser(id int) (string, error) {
// 	//returns the string resopnse , and errror response, if
// 	//you wanna ignore it then simply value, _ := findUser("uihfk203", 21) kardena

// }

//8:variadic functions:
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// 10: function as arg: (function type inside of function)
func execute(fn func()) {
	fn()
}

func main() {
	//1: basic function
	//done
	//2: parameters
	//done
	//3: return value :
	//done
	//4:
	fmt.Println(add(1, 2))
	//5:
	fmt.Println(getUserInfo())
	//6: ignoring values:
	name, _ := getUserInfo()
	fmt.Println(name)
	//7: done
	//8: variadic functions:
	fmt.Println(sum(1, 2, 3, 4, 5))

	//9: done
	//10: function as arg
	execute(func() {
		fmt.Println("hi from insides")
	})

	
}
