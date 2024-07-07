package main

import (
	"errors"
	"fmt"
)

type List interface {
	Add() ([]int, error)
}
type ListV1[T any] interface {
	// Add(idx int, t T)
	Append(t T) int
}
type AddList struct {
}

type User struct {
	Age  int
	Name string
}

func UseList() {
	var listv1 ListV1[int]
	myStruct := MyStruct{}
	listv1 = myStruct
	a := listv1.Append(12)
	fmt.Println(a)
}

type MyStruct struct {
	// 结构体的字段
}

// 实现 MyInterface 接口中的 Method1 方法
func (ms MyStruct) Append(idx int) int {
	// 实现 Method1 的逻辑
	return idx
}
func NewUser() {
	Users := User{}
	//指针
	User1 := &User{}
	//指针2
	User2 := new(User)
	//打印结构体
	fmt.Printf("%+v", Users)
	fmt.Printf("%+v", User1)
	fmt.Printf("%+v", User2)
}
func main() {
	// NewUser()
	// UseList()
	fmt.Println(Sum[int](12, 3, 2))
	fmt.Println(Max[int](12, 3, 2, 99))
}
func Sum[T Number](vals ...T) T {
	var t T
	for _, v := range vals {
		t = t + v
	}
	return t
}

// 泛形接口
type Number interface {
	//~ 引申类型
	~int | ~float32 | ~float64 | int8 | int16 | int32 | int64 | uint32
}

func Max[T Number](vals ...T) (T, error) {
	if len(vals) == 0 {
		var t T
		return t, errors.New("错误")
	}
	t := vals[0]
	for _, v := range vals {
		if v > t {
			t = v
		}
	}
	return t, nil
}
