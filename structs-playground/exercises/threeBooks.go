package main

import "fmt"

type Book struct {
	Title string
	Pages int
}

func threeBooks() {
	book1 := Book{
		Title: "DDIA",
		Pages: 200,
	}

	book2 := Book{
		Title: "WP7",
		Pages: 200,
	}

	book3 := Book{
		Title: "system design alex xu",
		Pages: 200,
	}

	fmt.Println(book1, book2, book3)

}
