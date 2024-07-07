package main

func sum(arr [5]int) (s int) {
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}
func main() {

}
