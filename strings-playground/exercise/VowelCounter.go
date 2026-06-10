package main

import ("fmt")
func countVowel(s string){
	count := 0 
	// length := len(s)
	for i:=0; i<len(s) ; i++ {
		
		// if s[i] == ('a' || 'e' || 'i' || 'o' || 'u'){
		// 	count++
		// }
		//above wont work because you would need to use s[i]='a' s[i]='e' s[i]="o". so you cannot just group everything into one
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u' : count ++
		}
	}
	fmt.Println(count)
}


