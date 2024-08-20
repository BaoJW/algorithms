package listNode

import "errors"

// 双链表

// NodeDouble 双链表节点
type NodeDouble[E any] struct {
	val  E
	prev *NodeDouble[E]
	next *NodeDouble[E]
}

// DoublyLinkedList 双链表实现
type DoublyLinkedList[E any] struct {
	head *NodeDouble[E]
	tail *NodeDouble[E]
	size int
}

func NewDoublyLinkedList[E any]() *DoublyLinkedList[E] {
	head := &NodeDouble[E]{}
	tail := &NodeDouble[E]{}
	head.next = tail
	tail.prev = head
	return &DoublyLinkedList[E]{head: head, tail: tail, size: 0}
}

// AddFirst 在链表头部添加元素
func (s *DoublyLinkedList[E]) AddFirst(e E) {
	newNode := &NodeDouble[E]{val: e}
	temp := s.head.next

	temp.prev = newNode
	newNode.next = temp
	s.head.next = newNode
	newNode.prev = s.head

	s.size++
}

// AddLast 在链表尾部添加元素
func (s *DoublyLinkedList[E]) AddLast(e E) {
	newNode := &NodeDouble[E]{val: e}
	temp := s.tail.prev

	temp.next = newNode
	newNode.prev = temp
	newNode.next = s.tail
	s.tail.prev = newNode

	s.size++
}

// AddAtIndex 在指定位置添加元素
func (s *DoublyLinkedList[E]) AddAtIndex(index int, element E) error {
	if index < 0 || index > s.size {
		return errors.New("index out of bounds")
	}
	if index == s.size {
		s.AddLast(element)
		return nil
	}

	p := s.getNode(index)
	newNode := &NodeDouble[E]{val: element}
	temp := p.prev

	p.prev = newNode
	temp.next = newNode
	newNode.prev = temp
	newNode.next = p

	s.size++
	return nil

}

// RemoveFirst 移除链表头部元素
func (s *DoublyLinkedList[E]) RemoveFirst() (E, error) {
	if s.size < 1 {
		return *new(E), errors.New("no such element")
	}

	removeNode := s.head.next
	temp := removeNode.next
	s.head.next = temp
	temp.prev = s.head

	// 删除节点置为nil
	removeNode.prev = nil
	removeNode.next = nil

	s.size--

	return removeNode.val, nil
}

// RemoveLast 移除链表尾部元素
func (s *DoublyLinkedList[E]) RemoveLast() (E, error) {
	if s.size < 1 {
		return *new(E), errors.New("no such element")
	}

	removeNode := s.tail.prev
	temp := removeNode.prev
	s.tail.prev = temp
	temp.next = s.tail

	// 删除节点置为nil
	removeNode.prev = nil
	removeNode.next = nil

	s.size--

	return removeNode.val, nil
}

// RemoveAtIndex 移除指定位置的元素
func (s *DoublyLinkedList[E]) RemoveAtIndex(index int) (E, error) {
	if index < 0 || index >= s.size {
		return *new(E), errors.New("index out of bounds")
	}
	removeNode := s.getNode(index)
	prev := removeNode.prev
	next := removeNode.next
	prev.next = next
	next.prev = prev

	// 删除节点置为nil
	removeNode.prev = nil
	removeNode.next = nil

	s.size--

	return removeNode.val, nil

}

// GetAtIndex 获取指定位置的元素
func (s *DoublyLinkedList[E]) GetAtIndex(index int) (E, error) {
	if index < 0 || index >= s.size {
		return *new(E), errors.New("index out of bounds")
	}
	p := s.getNode(index)
	return p.val, nil
}

// GetFirst 获取链表头部元素
func (s *DoublyLinkedList[E]) GetFirst() (E, error) {
	if s.size < 1 {
		return *new(E), errors.New("no such element")
	}
	return s.head.next.val, nil
}

// GetLast 获取链表尾部元素
func (s *DoublyLinkedList[E]) GetLast() (E, error) {
	if s.size < 1 {
		return *new(E), errors.New("no such element")
	}
	return s.tail.prev.val, nil
}

// Set 设置指定位置的元素
func (s *DoublyLinkedList[E]) Set(index int, val E) (E, error) {
	if index < 0 || index >= s.size {
		return *new(E), errors.New("index out of bounds")
	}
	p := s.getNode(index)
	oldVal := p.val
	p.val = val
	return oldVal, nil
}

// Size 获取链表大小
func (s *DoublyLinkedList[E]) Size() int {
	return s.size
}

// IsEmpty 检查链表是否为空
func (s *DoublyLinkedList[E]) IsEmpty() bool {
	return s.size == 0
}

// 获取指定索引的节点
func (s *DoublyLinkedList[E]) getNode(index int) *NodeDouble[E] {
	p := s.head.next
	for i := 0; i < index; i++ {
		p = p.next
	}
	return p
}
