package main 

func evenCounter(s []int) int{
	c:=0
	//given an array s this count even things
	for _, value := range s{
		if value%2==0 {c++}
	}
	return c
}
