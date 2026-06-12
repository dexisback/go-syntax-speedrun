package main 
import ("fmt")


func main(){
	//modify index of a slice
	nums:= [5]int{1,2,3,4,5}
	nums[2] = 100
	fmt.Println(nums)

	//get length of a slice toh len(nums)
	////iterating over a slice : done 
	//then it says ki bhai [5]int and [10]int are different, ofcourse they are. 
	var a [5]int
	var b [10]int
	//are not the same 

	_ = a
	_ = b ///yes i know these variables exist, and i am intentionally NOT using them 

	originalArray:=[]int{1,2,3}
	newArray := append(originalArray, 4)
	fmt.Println(newArray)

	//slicing a slice lol:
	 sliceToBeSliced := []int{1,2,3,4}
	 slicedSlice := (sliceToBeSliced[1:4])
	 fmt.Println(slicedSlice)


	 

}