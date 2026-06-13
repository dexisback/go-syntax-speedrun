func main
import ("fmt")


func groupWordsByLength(input []string) string{
	//takes in a slice of strings
	//goes through each index, counts the letters in that "value"
	//stores in a map (key should be int and value should be a slice (because multiple words can have the same lenght))
	groups:= make(map[[int]string])
	for _, word := range words {
		groups[len(word)] = append(groups[len(word)], word)
	}


}