package main1_test

import (
	main1 "gotest1111/basic/golangtest"
	"testing"
	"time"
)

func Test1(t *testing.T) {
	a := 1
	b := 2
	expe := 4
	actual := main1.Add(a, b)
	if expe != actual {
		t.Errorf("Add %d +%d=%d, expected %d", a, b, actual, expe)
	}
}
func Test2(t *testing.T) {
	t.Parallel()
	time.Sleep(1000)
	a := 1
	b := 2
	expe := 4
	actual := main1.Add(a, b)
	if expe != actual {
		t.Errorf("Add %d +%d=%d, expected %d", a, b, actual, expe)
	}
}
func Test3(t *testing.T) {
	t.Parallel()
	time.Sleep(1000)
	a := 1
	b := 2
	expe := 4

	actual := main1.Add(a, b)
	if expe != actual {
		t.Errorf("Add %d +%d=%d, expected %d", a, b, actual, expe)
	}
}
func BenchmarkXxx(t *testing.B) {
	a := 1
	b := 2
	expe := 3
	actual := main1.Add(a, b)
	if expe != actual {
		t.Errorf("Add %d +%d=%d, expected %d", a, b, actual, expe)
	}
}
func BenchmarkXxx1(t *testing.B) {

	expe := 3
	actual := main1.Add(1, 2)
	if expe != actual {
		t.Errorf("Add %d +%d=%d, expected %d", 1, 2, actual, expe)
	}
}
func ExampleXxx() {
	main1.PrintAdd(2, 1)
	// Unordered output:2;

}
func ExampleXxxPrintString() {
	main1.PrintString("  2", "1")
	// Unordered output:
	//           1
	// 2

}

func TestSub(t *testing.T) {
	t.Run("Test1", Test1)
	t.Run("Test2", Test2)
	t.Run("Test3", Test3)
}
