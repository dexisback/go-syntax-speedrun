package main

// import ("fmt")

func average(s []int) float64 {
	//takes a slice, iterates through it and then sabka sum / lenth of slice
	sum := 0
	for _, value := range s {
		sum += value
	}
	averagee := float64(sum) / float64(len(s))
	return averagee
}
