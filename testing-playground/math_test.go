package main

import "testing" //this testing package is used inside of test files because *testing.T is only needed there!

func addTest(t *testing.T) {
	result := adding(1, 2)
	want := 3
	if result != want {
		t.Errorf("expected 5, got %d", result)
	}
}
