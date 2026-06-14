package main

import "fmt"

type Book struct {
	Title string
	Pages int
}

// 1:
func (b Book) isLongBook() bool {

	return b.Pages > 300
}

// 2:
type Product struct {
	Name  string
	Price int
}

func (p Product) isExpensive() bool {
	return p.Price > 1000
}

// 4:
type BankAccount struct {
	Balance int
}

// increaser of balance:
func (b BankAccount) increaser() int {
	return b.Balance + 1000

}
func main() {
	book := Book{
		Title: "DDIA",
		Pages: 400,
	}

	fmt.Println(book.isLongBook())
	product := Product{
		Name:  "winkies cake",
		Price: 1000,
	}

	fmt.Println(product.isExpensive())

	curBal := BankAccount{
		Balance: 9000,
	}
	fmt.Println(curBal.increaser())
}
