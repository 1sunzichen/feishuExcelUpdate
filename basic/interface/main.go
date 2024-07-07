package main

import (
	"fmt"
	"unsafe"
)

type Car interface {
	Drive()
}
type Truck struct {
	Model string
}
type Tr interface {
	Drive()
}
type num1 struct {
	n1 int8  //1 被1整除
	n2 int16 //2 被2整除
}
type num2 struct {
	n1 int64 //8 被8整除
	n2 int64 //8
	n3 int64 //8
}

func (t Truck) Drive() {
	fmt.Println(t.Model)
}
func main() {
	var c Car = Truck{"t"}
	a := c.(Tr)

	fmt.Println(a, unsafe.Sizeof(a), unsafe.Sizeof(c))
	//内存对齐 ,对齐系数，长度
	fmt.Println(unsafe.Alignof(num1{}), unsafe.Sizeof(num1{}), unsafe.Alignof(num2{}), unsafe.Sizeof(num2{}))
}
