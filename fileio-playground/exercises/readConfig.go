package main

import (
	"errors"
	"fmt"
	"os"
)

func readConfig(path string) (string, error) {
	//file missing -> error
	//file empty -> error
	//baaki sab me content
	fileData, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	if len(fileData) == 0 {
		return "", errors.New("config file is empty")
	}

	return string(fileData), nil
}

func main() {
	content, err := readConfig("fullFile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(content)
}
