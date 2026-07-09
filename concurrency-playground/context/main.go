package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// create http request:
	req, err := http.NewRequest(http.MethodGet, "https://picsum.photos/200/300", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(req)
	//perform http request:

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	//else
	defer res.Body.Close()

	//get data from http response
	imageData, err := ioutil.ReadAll(res.Body)
	fmt.Println("downloaded image of size", len(imageData))
}
