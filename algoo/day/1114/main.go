package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"unicode"
)

// 搜索   数组中2个数字和等于target 返回2个值的索引
func main() {

	test5([]int{0, 0, 0, 0, 2, 3, 4, 1})
	// fmt.Println(twoSumHashTable([]int{1, 2, 3, 4}, 6))
}
func twoSumHashTable(nums []int, target int) []int {
	hashtable := map[int]int{}
	for k, v := range nums {
		if idx, ok := hashtable[target-v]; ok {
			return []int{idx, k}
		}
		hashtable[v] = k
	}
	return nil
}
func test2() {

	var magicConstant int
	fmt.Scan(&magicConstant)
	A := strconv.Itoa(int(magicConstant) - 20)
	B := strconv.Itoa(int(magicConstant) - 21)
	C := strconv.Itoa(int(magicConstant) - 18)
	D := strconv.Itoa(int(magicConstant) - 19)

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(A + ",1,12,7") // Write answer to stdout
	fmt.Println("11,8," + B + ",2")
	fmt.Println("5,10,3," + C)
	fmt.Println("4," + D + ",6,9")
}
func test3(N int) {
	// var N int
	// fmt.Scan(&N)
	sum := 1
	for i := N; i > 0; i-- {
		sum *= i
	}
	sumzero := 0
	for sum%10 == 0 {
		sum /= 10
		sumzero += 1
	}
	if N == 0 {
		sumzero = 1
	}
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(sumzero) // Write answer to stdout
}
func test4() {
	var str = "S0, 1f u turn 0n u'r br4in u wi11 underst4nd"
	var s float64
	var n float64
	// 使用range关键字遍历字符串
	for _, char := range str {
		// fmt.Printf("%c ", char)
		if unicode.IsDigit(char) {
			n++
		} else if unicode.IsLetter(char) {
			s++
		}
	}
	fmt.Println(s / n)
	// fmt.Fprintln(os.Stderr, "Debug messages...")
	fmt.Println(math.Round(float64(s / n))) // Write answer to stdout
}

type Pair struct {
	Key   int
	Value int
}

func test5(mountainH []int) {
	for {
		H := make(map[int]int, 8)
		for i := 0; i < len(mountainH); i++ {
			// mountainH: represents the height of one mountain.
			// var mountainH int
			// fmt.Scan(&mountainH)
			H[i] = mountainH[i]
		}
		pairs := make([]Pair, 0, len(H))
		for key, value := range H {
			pairs = append(pairs, Pair{key, value})
		}
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].Value > pairs[j].Value
		})
		for _, pair := range pairs {
			fmt.Print(pair.Key)
			// fmt.Printf("学生：%s，成绩：%d\n", pair.Key, pair.Value)
		}
		// fmt.Fprintln(os.Stderr, "Debug messages...")
		// The index of the mountain to fire on.

	}
}
