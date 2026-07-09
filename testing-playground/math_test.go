package main

//single testing

// func TestAdd(t *testing.T) {
// 	result := adding(1, 2)
// 	want := 3
// 	if result != want {
// 		t.Errorf("expected %d, got %d", want, result)
// 	}
// }

// multiple:

import "testing"

func TestAdd(t *testing.T) {
	if adding(2, 3) != 5 {
		t.Error(" 2+ 3 failed")
	} else {
		t.Log("success. 2+3 ✅")
	}

	if adding(10, 5) != 15 {
		t.Error(" 10 + 5 failed 🥀")
	} else {
		t.Log("success. 10 +5 ✅")
	}
}

//table driven tests (more common in go actually)

func TestAddTableDriven(t *testing.T) {
	//pehle toh type bana lo
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{2, 3, 5},
		{10, 5, 15},
		{6, 7, 13},
	}

	for _, tt := range tests {
		testKaResult := adding(tt.a, tt.b)
		if testKaResult != tt.expected {
			t.Error("test", tt.a, "+", tt.b, "failed🥀")
		} else {
			t.Log("test", tt.a, "+", tt.b, "success ✅")
		}

	}

}



//phase 7: testing errors


func TestDivide(t *testing.T){
	_, err := dividing(5,0)

	if err == nil {
		t.Errorf("denominator 0 hai 🥀")
	}
} 


//benchmarking : 

//function name should start with "Benchmark" and parameter b *testing.B hojata
//why n? because YOU DONT choose how many appropriate times to run it. go appropriately figures out kitne time lagega and then utna woh chala ke measure karta hai 


func BenchmarkAdd(b *testing.B){
	for i:= 0; i<b.N ; i++{
		adding(2,3)
	}
}

//coverage: percentage dedeti kitne tests pass hore 

//testing file i/o 
