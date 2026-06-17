package main

import (
	"fmt"
	"os"
)

//read entire file:

func main() {
	//phase 1: read entire file notes.txt
	// data, err := os.ReadFile("files/notes.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// //else:
	// fmt.Println(string(data))
	// //phase 2: read missing file:

	// nonExistenData, err := os.ReadFile("nonExistent.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// //else
	// fmt.Println(nonExistenData)

	// //phase 3: empty file validation

	// emptyFileData, err := os.ReadFile("files/emptyNotes.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// //else:(done this already)
	// if len(emptyFileData) == 0 {
	// 	errors.New("file is empty")
	// 	return
	// }

	//phase 4: write file:
	// err := os.WriteFile("files/notes.txt",
	// 	[]byte("hellow this is me modifying for the first time"),
	// 	0644,
	// )

	// if err != nil {

	// 	return
	// }
	// //else:
	// fmt.Println("success")
	//phase 5: read after write
	//write os.ReadFile after os.WriteFile (write file only returns one error, if there was any error while writing the thang)

	// //phase 6: os.WriteFile always rewrites, how do you append? using os.O_APPEND
	// file, error := os.OpenFile("files/notes.txt",
	// 	os.O_APPEND|os.O_WRONLY,
	// 	0644,
	// )

	// if error != nil {
	// 	fmt.Println(error)
	// 	return
	// }
	// defer file.Close()
	// file.WriteString("\n new line")

	// //create file. using os.Create()
	// file, err := os.Create("files/createdThroughCode.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(file)

	// //phase 9: file handling
	// 	//create file. using os.Create()
	// file, err := os.Create("files/createdThroughCode.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(file)
	// //isme file returns *os.File this is connection to the file  , like db connection but ok not needed right now
	//phase 10: always close the thang

	//by using defer file.Close() (kinda like finally{})
	// //phase 11: read line by line :
	// file, err := os.Open("files/linebyline.txt")

	// if err != nil {
	// 	return
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// } //used for logs, csv , large files, etc

	// //os.ReadFile loads everything, but scanner.Scan() scans and reads line by line, so load nahi padega systum pe if large file read hori ho

	// filess, errr := os.ReadDir("files")

	// if errr != nil {
	// 	return
	// }

	// for _, file := range filess {
	// 	fmt.Println(file.Name)
	// 	fmt.Println(file.IsDir())
	// 	fmt.Println(file.Type())
	// 	fmt.Println(file.Type().IsDir())

	// }

	// //phase 14: check existence:
	// _, err := os.Stat("files/nootes.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("success")

	//phase 15: delete file
	err := os.Remove("files/notes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//else
	fmt.Println("succes")

	//you have to read configs and then unmarhsall them for bes mostly
}
