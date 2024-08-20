package queue

import (
	"personal/algorithms/dataStructures/listNode"
)

// 单链表实现消息队列

type MySinglyListStack[E any] struct {
	list *listNode.SinglyLinkedList[E]
}

func NewMySinglyListStack[E any]() *MySinglyListStack[E] {
	return &MySinglyListStack[E]{
		list: listNode.NewSinglyLinkedList[E](),
	}
}

// Push 向队尾插入元素，时间复杂度O(1)
func (m *MySinglyListStack[E]) Push(e E) {
	m.list.AddLast(e)
}

// Pop 从队头删除元素，时间复杂度O(1)
func (m *MySinglyListStack[E]) Pop() (E, error) {
	return m.list.RemoveFirst()
}

// Peek 查看队头元素，时间复杂度O(1)
func (m *MySinglyListStack[E]) Peek() (E, error) {
	return m.list.GetFirst()
}

func (m *MySinglyListStack[E]) Size() int {
	return m.list.Size()
}
