package main 


//nums := []int{5,9,2,7,1}
//isme se find largest, so iterate and find the largest value


func findLargest(s []int) int {
	largest := s[0] 
	for _, value := range s {
		if value> largest {
			//update largest 
			largest = value 
		}
	}
	return largest 
}