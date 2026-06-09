// package main

// import "fmt"

// func main(){
// 	name:="amaan"
// 	fmt.Println(name)
// }

//--

//length
// func main(){
// 	name:="Amaan"
// 	fmt.Println(name)
// 	fmt.Println(len(name))
// }

//----

//indexing:
// func main(){
// 	name:="Amaan"
// 	fmt.Println(name[0])
// 	fmt.Println(name[1])
// 	fmt.Println(name[2])
// }
//returns 65 because a go string is a sequence of bytes. so the value of the byte A it returns
//if you want the chharacter
// func main(){
// 	name:="Ayaan"
// 	fmt.Println("%c\n", name[0]) //prints A
// 	//or:
// 	fmt.Println(string(name[0])) //prints A
// 	fmt.Println(string(name)) //prints Ayaan
// }

//phase 4: bytes:

// package main
// import "fmt"

// func main(){
// 	name:="Amaan"
// 	fmt.Printf("%T\n", name[0])
// }

// package main
// import "fmt"

// func main(){
// 	ch:="A"
// 	fmt.Println(ch)
// 	fmt.Printf("%T\n", ch)
// }

// package main
// import "fmt"

// func main(){
// 	word:="namaste"
// 	fmt.Println(word)
// 	fmt.Println(len(word))
// }

//--------

// //phase 7: range:
// package main
// import "fmt"
// func main(){
// 	word:="Amaanikaaa"
// 	for _, ch:= range word {
// 		fmt.Println(ch)
// 	}
// }

// package main
// import "fmt"

// func main(){
// fmt.Println("hellow")
// }

//printing actual characters with range
// package main
// import "fmt"

// func main(){
// 	variable_name:="Amaan"
// 	for _, ch := range variable_name{
// 	fmt.Printf("%c", ch)}
// }

//strings package (the default one we talked about)

// package main
// import (
// 	"fmt"
// 	// "strings"
// )
// var pizza = "I wanna eat pizza"

// func testfunction(){
// 	fmt.Println("hellow we are testing compiler will ignore variable `pizza` or not")
// }

// func main(){
// 	variable_name :="Ayaan"
// 	//even though we wont use this variable, but i will silence it so the compiler does NOT throw in errors (workaruond)
// 	_ = variable_name
// 	email := "amaangmail.com"
// 	_ = email
// 	// second_variable := "a"
// 	// fmt.Println(strings.ToUpper(variable_name))
// 	// fmt.Println(strings.ToLower(variable_name))
// 	// fmt.Println(strings.Count(variable_name, second_variable))
// 	// fmt.Println(strings.CutSuffix(variable_name, second_variable))
// 	// contains_value := strings.Contains(email, "@")
// 	// fmt.Println(contains_value)
// 	testfunction()
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"unicode"
// )

// func capitaliser(s string) string {
// 	if len(s) == 0 {
// 		return s
// 	}
// 	runes := []rune(s)
// 	runes[0] = unicode.ToUpper(runes[0])
// 	return string(runes)

// }

// func main() {
// 	token := "bearer abc123"
// 	//now i need to capitalise the first letter . do this using the unicode package
// 	token = capitaliser(token)

// 	fmt.Println(strings.HasPrefix(token, "Bearer"))
// }




//join, replace, trim, builder
//and mini exercise 1/2/3/4 do 


// package main 
// import ("fmt"
// "strings")


// func main(){
// 	path:= "/usr/home/bin/my_stuff"
// 	parts:= strings.Split(path, "/")
// 	fmt.Println(parts)
// }



// //join

// package main
// import(
// 	"fmt"
// 	"strings"
// )

// func main(){
// 	parts:=[]string{"user", "profile", "settings"}
// 	result := strings.Join(parts, "/")
// 	fmt.Println(result)
// }


//