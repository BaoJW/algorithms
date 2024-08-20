package stack

import "personal/algorithms/dataStructures/listNode"

// MyDoublyListStack 双链表实现栈
type MyDoublyListStack[E any] struct {
	list *listNode.DoublyLinkedList[E]
}

func NewMyDoublyListStack[E any]() *MyDoublyListStack[E] {
	return &MyDoublyListStack[E]{
		list: listNode.NewDoublyLinkedList[E](),
	}
}

// Push 向栈顶(队尾)插入元素，时间复杂度O(1)
func (m *MyDoublyListStack[E]) Push(e E) {
	m.list.AddLast(e)
}

// Pop 从栈顶(队尾)删除元素，双链表时间复杂度O(1)
func (m *MyDoublyListStack[E]) Pop() (E, error) {
	return m.list.RemoveLast()
}

// Peek 查看栈顶(队尾)元素，双链表时间复杂度O(1)
func (m *MyDoublyListStack[E]) Peek() (E, error) {
	return m.list.GetLast()
}

func (m *MyDoublyListStack[E]) Size() int {
	return m.list.Size()
}
