package main 

import ("fmt")



func palindromeChecker(s string) {
	newWord:=""
	for i:= len(s)-1 ;i>=0 ; i--{
		newWord += string(s[i])
	}

	if newWord == s {
		fmt.Println("yes")
	} else{
		fmt.Println("no")
	}
}