package main

import (
	"fmt"
	"reflect"
)

func Add(a, b int) int {
	return a + b
}
func Add2(a, b int) int {
	return a - b
}

// 反射调用方法
func CallAdd(fn func(a int, b int) int) {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return
	}
	//入参
	argv := make([]reflect.Value, 2)
	argv[0] = reflect.ValueOf(1)
	argv[1] = reflect.ValueOf(2)
	res := v.Call(argv)
	//处理返回参数
	fmt.Println(res[0].Int())
}
func main() {
	CallAdd(Add2)
	var a float32 = 2.11
	fmt.Printf("%v", reflect.TypeOf(a))
	var k = reflect.ValueOf(a).Kind()
	fmt.Println(k)
}
