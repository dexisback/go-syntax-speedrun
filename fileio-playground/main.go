package main

import (
	"errors"
	"fmt"
	"os"
)

//read entire file:

func main() {
	//phase 1: read entire file notes.txt
	data, err := os.ReadFile("files/notes.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	//else:
	fmt.Println(string(data))
	//phase 2: read missing file:

	nonExistenData, err := os.ReadFile("nonExistent.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//else
	fmt.Println(nonExistenData)

	//phase 3: empty file validation

	emptyFileData, err := os.ReadFile("files/emptyNotes.txt")
	if err != nil {
		fmt.Println(err)
	}
	//else:
	if len(emptyFileData) == 0 {
		errors.New("file is empty")
		return
	}

}
