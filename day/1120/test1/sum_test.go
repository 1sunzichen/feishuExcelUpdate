package main

import "testing"

func TestSum(t *testing.T) {
	number := [5]int{1, 2, 2, 4, 5}
	got := sum(number)
	want := 14
	if got != want {
		t.Errorf("got %d want %d given,%v", got, want, number)
	}
}
