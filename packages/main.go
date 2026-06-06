package main

import (
	"fmt"
	"packages/auth"
)


func main(){
	token:= auth.GenerateToken()
	hashedPassword := auth.HashPassword("secret123")
	IsAuthenticated := auth.IsAuthenticated()


	fmt.Println(token)
	fmt.Println(hashedPassword)
	fmt.Println(IsAuthenticated)
}