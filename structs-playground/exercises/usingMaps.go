package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

// create a map, map[int]Product
func usingMaps() {
	//to make maps:
	//first make a map using the make() function, and specify its types
	//so like ages:= make(map[int]string)
	//then add values onto it. so like: ages["Alice"] = 28

	inputMap := make(map[int]Product)
	//now you can add values inputMap[134] = {Name: Britania cake, Price: 20}  //the 134 being its id and the value being a type(struct) Product
	inputMap[134] = Product{Name: "britania cake", Price: 20}
	inputMap[135] = Product{Name: "winkies cake", Price: 30} //because premium
	product1, exists1 := inputMap[134]
	product2, exists2 := inputMap[135]

	if exists1 {
		fmt.Println("product 1 found ->", product1)
	} else {
		fmt.Println("not found")
	}

	if exists2 {
		fmt.Println("product 2 found ->", product2)
	} else {
		fmt.Println("not found")
	}

}
