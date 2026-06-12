//reverse and append

package main

func reverseSlice(s []int) []int{
	//takes in an array of integers, and:
	//appends them into a new array
	nA := []int{}
	for i := len(s) - 1; i >= 0; i-- {
		nA = append(nA, s[i])
	}
	return nA
}
