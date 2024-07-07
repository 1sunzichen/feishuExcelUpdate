package main

var FruitColor map[string]string

func AddFruit(name, color string) {
	FruitColor[name] = color
}

var StudentScore map[string]int

func GetScore(name string) int {
	score := StudentScore[name]
	return score
}
func main() {

	AddFruit("string", "red")
	GetScore("string")

}

//向值为nil的map添加元素会触发panic   函数中应先判断map的值是否是nil

//查询map时应判断元素是否存在
