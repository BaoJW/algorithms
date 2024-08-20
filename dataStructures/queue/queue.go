package queue

// MyQueue 队列：只能在一端插入元素，另一端删除元素
type MyQueue[E any] interface {
	Push(e E)         // 向队尾插入元素，时间复杂度O(1)
	Pop() (E, error)  // 从队头删除元素，时间复杂度O(1)
	Peek() (E, error) // 查看队头元素，时间复杂度O(1)
	Size() int        // 返回队列中的元素个数，时间复杂度O(1)
}
