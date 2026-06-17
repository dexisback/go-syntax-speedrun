package main

import (
	"bufio"
	"fmt"
	"os"
)

func longestLine(path string) (int, error) {
	//read a file
	//buffer+scanner
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	longestLine := ""

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > len(longestLine) {
			//then update longestLine
			longestLine = line
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error reading file", err)
		return 0, err
	}

	fmt.Println("longest line", longestLine)
	return len(longestLine), nil
}



func main(){
	longestLine("fullFile.txt")
}