package stack

import "personal/algorithms/dataStructures/array"

// 数组实现栈
// 本质上就是通过环形数组实现

type MyArrayStack[E any] struct {
	arr *array.CycleArray[E]
}

func NewMyArrayStack[E any]() *MyArrayStack[E] {
	return &MyArrayStack[E]{
		arr: array.NewCycleArray[E](),
	}
}

// Push 向栈顶(队尾)插入元素，时间复杂度O(1)
func (m *MyArrayStack[E]) Push(e E) {
	m.arr.AddLast(e)
}

// Pop 从栈顶(队尾)删除元素，时间复杂度O(1)
func (m *MyArrayStack[E]) Pop() (E, error) {
	return m.arr.RemoveFirst()
}

// Peek 查看栈顶(队尾)元素，时间复杂度O(1)
func (m *MyArrayStack[E]) Peek() (E, error) {
	return m.arr.GetFirst()
}

func (m *MyArrayStack[E]) Size() int {
	return m.arr.Size()
}
