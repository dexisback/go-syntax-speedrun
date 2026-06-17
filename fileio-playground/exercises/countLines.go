package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(path string) (int, error) {
	//takes in a path of a file , and does count lines
	//go ki os library does NOT have a default counting line feature
	//so we pair it with bufio scanner library and use a counter
	file, err := os.Open(path)

	if err != nil {
		return 0, err
	}
	//else:
	defer file.Close()
	lineCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineCount++
	}

	//check for errors during scanning:
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}

func main() {
	count, err := countLines("users.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Total lines are", count)
}
