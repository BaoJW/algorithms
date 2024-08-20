package queue

import "personal/algorithms/dataStructures/array"

// 数组实现队列
// 本质上就是通过环形数组实现

type MyArrayQueue[E any] struct {
	arr *array.CycleArray[E]
}

func NewMyArrayQueue[E any]() *MyArrayQueue[E] {
	return &MyArrayQueue[E]{
		arr: array.NewCycleArray[E](),
	}
}

// Push 向队尾插入元素，时间复杂度O(1)
func (m *MyArrayQueue[E]) Push(e E) {
	m.arr.AddLast(e)
}

// Pop 从队头删除元素，时间复杂度O(1)
func (m *MyArrayQueue[E]) Pop() (E, error) {
	return m.arr.RemoveFirst()
}

// Peek 查看队头元素，时间复杂度O(1)
func (m *MyArrayQueue[E]) Peek() (E, error) {
	return m.arr.GetFirst()
}

func (m *MyArrayQueue[E]) Size() int {
	return m.arr.Size()
}
