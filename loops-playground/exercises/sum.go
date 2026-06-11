package main 
//input: nums := []int{1,2,3,4,5}
//output = 15

//sum hora har index ka:

import ("fmt")


func summer(s []int){
	sum:= 0
	// inputArray:= []int{1,2,3,4,5}
	//iterate over everything 
	
	for _, value := range s {
		sum = value+sum
	}
	fmt.Println(sum)
}



