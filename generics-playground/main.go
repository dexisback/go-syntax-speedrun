package main

import ("fmt"
"golang.org/x/exp/constraints")

func Adder(a int, b int) int {
	return a + b
}

func main() {
	result := Adder(1, 2)
	fmt.Println(result)
	resultWithGenerics_zero := AddNew(1.1, 2.2)
	resultWithGenerics_one := AddNew(1, 2)
	fmt.Println(resultWithGenerics_zero)
	fmt.Println(resultWithGenerics_one)
	resultSimpleWithInterface := AddNewSimple(1, 2.2)
	fmt.Println(resultSimpleWithInterface)

}

//the above is great and all, but what if we wanna pass in something OTHER than int
//then you will not get the answer with the same function
//what you will have to do in a prod system is make an Addint and an AddFloat
//two functions, returning different types, and you call each when needed based on the situation

func AddFloat(a float64, b float64) float64 {
	return a + b
}

//but obv repetitive
//now with generics you can do that

// we are gonna write an "AddNew function which will be a ONE STOP solution to everything "
func AddNew[T int | float64](a T, b T) T {
	return a + b
}

//this is cool, but the next bottleneck is ki there are a lot of things which we wrote
//like T int | float 4
//instead, make an interface for that data type

type Num interface {
	int | int8 | int16 | float64
}

func AddNewSimple[T Num](a T, b T) T {
	return a + b
}



//now, this is still a lot of verbose
//what we can do is , get rid of Num and the interface you made 
//and type in constraints.Ordered by importing "golang.org/x/exp/constraints


//this library take cares of the long work we needed to do 
 

func AddWithLibrary[T constraints.Ordered](a T, b T)T{
	return a + b 
}  

//now this addwithblibrary takes care and a generic of EVERY possible data type (NUMERICAL). so like int 64, int 32, whatever fuckery is out there

//but this would involve everything 
//if you want lets say only two you make your type

type UserID int //something like this

//----

//real world examples: map:
//input : [1,2,3], (n) => n * 2
//output  [2,4,6]

func MapValues(values []int, mapFunc func){

}
