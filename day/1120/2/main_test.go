package main

import "testing"

func TestMain(t *testing.T) {
	i := 2
	j := 1
	want := i
	got := If(i > j, i, j)
	if got != want {
		t.Errorf("want %v,got %v", want, got)
	}
}
