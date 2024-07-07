package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "10000"
	ret1, _ := strconv.ParseInt(str, 10, 64)

	fmt.Println(ret1)
	ret2, _ := strconv.Atoi(str)
	fmt.Println(ret2)
}
