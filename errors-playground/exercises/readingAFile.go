// read two files, (one will be full and one will be empty), generalise reader, and when any file is empty return error for it
package main

import (
	"fmt"
	"os"
)

func ReadFile(path string) (string, error) {
	//takes in a path and returns the content if file has content, and returns an error if it is empty
	data, err := os.ReadFile(path)

	if err != nil {
		//mtlb error hai
		return "", err
	}

	if len(data) == 0 {
		return "", fmt.Errorf("file is empty")
	}

	//else if everything is fine, return the content
	return string(data), nil

}

func execute() {
	content, err := ReadFile("emptyFile.txt")
	if err != nil {
		//then print the error
		fmt.Println("error is ", err)
		return
	}
	//else print the content
	fmt.Println(content)
}

// func main() {
// 	execute()
// }

//NOTE: go treats an empty file as a succesfull read, so data is empty but error will be nill
//to truly make a file empty, you need to do the thing i did above which is

/*if len(data) == 0 {
	return "", fmt.Errorf("file is empty")
}*/
