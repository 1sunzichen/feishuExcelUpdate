package main

// defer的用法
import "fmt"

func main() {
	fmt.Println(f1(2))
}

func f(x int) (r int) {
	defer func() {
		r += x
	}()
	return x + x
}
func f1(x int) (r int) {
	defer func(p int) {
		fmt.Printf("r: %d,p: %d,x: %d\n", r, p, x)
		r = p + r
	}(r)
	return x + x
}
func f10(x int) (r int) {
	defer func(p int) {
		fmt.Printf("r: %d,p: %d,x: %d\n", r, p, x)
		r = p + x
	}(r)
	return x + x
}
func f2(x int) (r int) {
	defer func(p int) {
		fmt.Printf("r: %d,p: %d,x: %d\n", r, p, x)
		r = r + p
	}(x)
	return x + x
}
