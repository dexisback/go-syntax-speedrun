//playground 

package main
import (
	"fmt"
	"go/auth"
)

func main(){
	token:= auth.GenerateToken()
	hashed:= auth.HashPassword("secret")
	isAuth:= auth.IsAuthenticated()


	fmt.Println(token)
	fmt.Println(hashed)
	fmt.Println(isAuth)

}