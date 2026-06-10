//slice: temperory array?
//package main

//import( "fmt" )

//func main(){
//	fruits := []string{"apple", "oranges"}
//	fmt.Println(fruits)

//}


//make function: used to allocate empty slice with pre-defined length and an optional capacity
//func main(){
//	numbers:= make([]int, 3,5)
//}


//slicing an array:
package main
import ("fmt")
func main(){


arr:= [5]int{10,20,30,40,50}
subSlice:= arr[1:4]
fmt.Println(subSlice)
}


