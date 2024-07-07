package main

import "fmt"

type People struct {
	Name string
	age  int
}

func (p *People) SetName(name string) {
	p.Name = name
}
func (p *People) SetAge(age int) {
	p.age = age
}
func (p *People) GetAge() int {
	return p.age
}
func main() {
	p := &People{age: 11, Name: "111"}
	fmt.Println(p.GetAge())
	p.SetAge(1222)
	fmt.Println(p.GetAge(), p.age)

}
