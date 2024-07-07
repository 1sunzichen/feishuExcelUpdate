package main

import (
	"context"
	"fmt"
)

func main() {
	con := context.WithValue(context.Background(), "12", 12)
	fmt.Println(con.Done())
}
