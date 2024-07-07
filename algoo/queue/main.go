package main

type arrayQueue struct {
	nums        []int //用来存储队列的元素
	front       int   //队首指针 指向队首元素
	queSize     int   //队列长度
	queCapacity int   //队列容量
}

func newArrayQueue(queCapacity int) *arrayQueue {
	return &arrayQueue{
		nums:        make([]int, queCapacity),
		queCapacity: queCapacity,
		front:       0,
		queSize:     0,
	}
}

/* 入队 要放在最后一个*/
func (q *arrayQueue) push(num int) {
	// 当 rear == queCapacity 表示队列已满
	if q.queSize == q.queCapacity {
		return
	}
	rear := (q.front + q.queSize) % q.queCapacity
	q.nums[rear] = num
	q.queSize++
}

/* 出队 */
func (q *arrayQueue) pop() any {
	num := q.peek()
	// 队首指针向后移动一位，若越过尾部则返回到数组头部
	q.front = (q.front + 1) % q.queCapacity
	q.queSize--
	return num
}

/* 访问队首元素 */
func (q *arrayQueue) peek() any {
	if q.isEmpty() {
		return nil
	}
	return q.nums[q.front]
}

/* 判断队列是否为空 */
func (q *arrayQueue) isEmpty() bool {
	return q.queSize == 0
}

/* 获取 Slice 用于打印 */
func (q *arrayQueue) toSlice() []int {
	rear := (q.front + q.queSize)
	if rear >= q.queCapacity {
		rear %= q.queCapacity
		return append(q.nums[q.front:], q.nums[:rear]...)
	}
	return q.nums[q.front:rear]
}
func main() {

}
