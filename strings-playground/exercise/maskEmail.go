package main 




// func maskEmail(email string) string{
// 	//you have to iterate from index 1 till @ is received
// 	newEmail := ""

// 	for i:=1 ; i<len(email); i++ {
// 		if email[i]=='@'{
// 			break
// 		}else{
// 		  newEmail := string(email[0])+"*"
// 		}

// 	}

// }


func maskEmail(email string) string {

newEmail := string(email[0])

for i:=1 ; i< len(email) ; i++ {
	if email[i] == '@' {
		newEmail += email[i:]
		break
	} else{
		newEmail += "*"
	}
}
return newEmail

}