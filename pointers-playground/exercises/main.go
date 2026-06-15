package main

import "fmt"

// 1:
func main() {
	x := 5
	ptr := &x

	fmt.Println(*ptr)
	fmt.Println(x)

	*ptr = 20
	//now x becomes 20
	fmt.Println(x, "second wala")

	account := BankAccuont{
		Bal: 300,
	}

	fmt.Println(account.Deposit(500))
	alpha := 25
	fmt.Println(alpha, "is the old value before reset")

	Reset(&alpha)
	fmt.Println(alpha, "is the new value")

}

//2: implement func Double

// if we had :
func Double(num int) {
	num = num * 2
}

//then with the above a copy gets dobled, not the original thing

func Doble(num *int) { //w this we are receiving the address of the integer, "go to the address and get the actual value"
	*num = *num * 2
}

//now when i will call this function i would need to call Double(&x) and not just Double(x)

//3:implement resetter
func Reset(counter *int) {
	*counter = 0
}

//now you call it like:

//set its value to 0

type BankAccuont struct {
	Bal int
}

//method

func (b *BankAccuont) Deposit(amount int) int {
	return b.Bal + amount
}

type User struct {
	Name string
}

// method:
func (u *User) Rename(newName string) string {
	u.Name = newName
	return u.Name
}
