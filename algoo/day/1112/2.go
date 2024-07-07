package main

// https://leetcode.cn/problems/implement-stack-using-queues/
// 队列实现栈
type MyStack struct {
	data []int
}

func Constructor() MyStack {
	return MyStack{
		// 设置栈的长度为 0，容量为 16
		data: make([]int, 0, 16),
	}
}

func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyStack) Pop() int {
	top := this.Top()
	this.data = this.data[:len(this.data)-1]
	return top
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	val := this.data[len(this.data)-1]
	return val
}

func (this *MyStack) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
