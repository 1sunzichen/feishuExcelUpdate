package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Score int
}
type ByScore []Student

func (s ByScore) Len() int {
	return len(s)
}
func (s ByScore) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByScore) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}
func printStudnet(students []Student) {
	for _, student := range students {
		fmt.Printf("Name %s,Score%d \n", student.Name, student.Score)
	}
}
func partconfigurations(num []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && num[i] <= num[left] {
			i++
		}
		for i > j && num[j] >= num[left] {
			j--
		}
		num[i], num[j] = num[j], num[i]
	}
	num[i], num[left] = num[left], num[i]
	return i
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	p := partconfigurations(nums, left, right)
	quickSort(nums, left, p-1)
	quickSort(nums, p+1, right)
}
func main() {
	numberss := []int{86, 3, 1, 3, 55, 23, 12, 5, 3555, 6}
	quickSort(numberss, 0, len(numberss)-1)
	fmt.Println(numberss)
	// student := []Student{
	// 	{
	// 		Name:  "xiaoming",
	// 		Score: 12,
	// 	},
	// 	{
	// 		Name:  "xiaoming1",
	// 		Score: 11,
	// 	},
	// 	{
	// 		Name:  "xiaoming2",
	// 		Score: 13,
	// 	},
	// }
	// sort.Sort(ByScore(student))

	// fmt.Println("排序后")
	// printStudnet(student)
}
