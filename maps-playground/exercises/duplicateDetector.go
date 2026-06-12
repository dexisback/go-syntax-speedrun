package main
// import ("fmt")
func duplicateDetect(s []int) bool{
	//given an array you should be able to detect
	//if any number repeats in the array
	//so simple , we store and increase the counter . if any counter ever increases to 2 after increasing , that means repeated so then print the key and value
	trackMap:= make(map[int]int)
	for _,value := range s{
		//check if value alr exists in trackMap
		_, exists := trackMap[value]
		//if it does return true immediately
		if exists {
			return true
		} 
		trackMap[value]=1 ////if it doesnt then continue and add it to the map


		
	}
	return false //because beyond 1 value hume kuch kaam hi nahi

}



