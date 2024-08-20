package stack

import (
	"personal/algorithms/dataStructures/listNode"
)

// MySinglyListStack 单链表实现栈
type MySinglyListStack[E any] struct {
	list *listNode.SinglyLinkedList[E]
}

func NewMySinglyListStack[E any]() *MySinglyListStack[E] {
	return &MySinglyListStack[E]{
		list: listNode.NewSinglyLinkedList[E](),
	}
}

// Push 向栈顶(队尾)插入元素，时间复杂度O(1)
func (m *MySinglyListStack[E]) Push(e E) {
	m.list.AddLast(e)
}

// Pop 从栈顶(队尾)删除元素，单链表时间复杂度O(N)
func (m *MySinglyListStack[E]) Pop() (E, error) {
	return m.list.RemoveLast()
}

// Peek 查看栈顶(队尾)元素，单链表时间复杂度O(N)
func (m *MySinglyListStack[E]) Peek() (E, error) {
	return m.list.GetLast()
}

func (m *MySinglyListStack[E]) Size() int {
	return m.list.Size()
}
