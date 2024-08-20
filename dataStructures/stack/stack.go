package stack

// MyStack 栈：只能在某一端插入和删除元素
type MyStack[E any] interface {
	Push(e E)         // 向栈顶插入元素，时间复杂度O(1)
	Pop() (E, error)  // 从栈顶删除元素，时间复杂度O(1)
	Peek() (E, error) // 查看栈顶元素，时间复杂度O(1)
	Size() int        // 返回栈中的元素个数，时间复杂度O(1)
}
