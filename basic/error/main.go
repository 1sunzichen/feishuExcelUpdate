package main

import (
	"errors"
	"fmt"
)

func ex() {
	err := errors.New("this error")
	wrapError := fmt.Errorf("some context:%w", err)
	if _, ok := wrapError.(interface{ Unwrap() error }); ok {
		fmt.Println("wrap")
	}
}
func main() {
	ex()
}
