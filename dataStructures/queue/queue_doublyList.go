package queue

import (
	"personal/algorithms/dataStructures/listNode"
)

type MyDoublyListStack[E any] struct {
	list *listNode.DoublyLinkedList[E]
}

func NewMyDoublyListStack[E any]() *MyDoublyListStack[E] {
	return &MyDoublyListStack[E]{
		list: listNode.NewDoublyLinkedList[E](),
	}
}

// Push 向队尾插入元素，时间复杂度O(1)
func (m *MyDoublyListStack[E]) Push(e E) {
	m.list.AddLast(e)
}

// Pop 从队头删除元素，时间复杂度O(1)
func (m *MyDoublyListStack[E]) Pop() (E, error) {
	return m.list.RemoveFirst()
}

// Peek 查看队头元素，时间复杂度O(1)
func (m *MyDoublyListStack[E]) Peek() (E, error) {
	return m.list.GetFirst()
}

func (m *MyDoublyListStack[E]) Size() int {
	return m.list.Size()
}
