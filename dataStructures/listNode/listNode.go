package listNode

import (
	"errors"
	"fmt"
)

// 链表
// 一条链表并不需要一整块连续的内存空间存储元素。链表的元素可以分散在内存空间的天涯海角，通过每个节点上的 next, prev 指针，将零散的内存块串联起来形成一个链式结构
// 优点1：提高内存的利用效率，链表的节点不需要挨在一起，给点内存 new 出来一个节点就能用，操作系统会觉得这娃好养活
// 优点2：它的节点要用的时候就能接上，不用的时候拆掉就行了，从来不需要考虑扩缩容和数据搬移的问题，理论上讲，链表是没有容量限制的（除非把所有内存都占满，这不太可能）
// 局限性1：数组支持通过索引快速访问元素，而链表就不支持

// 链表实现关键点
// 关键点1: 同时持有头尾节点的引用
// 因为在软件开发中，在容器尾部添加元素是个非常高频的操作，双链表持有尾部节点的引用，就可以在O(1)的时间复杂度内完成尾部添加元素的操作。
// 对于单链表来说，持有尾部节点的引用也有优化效果。比如你要在单链表尾部添加元素，如果没有尾部节点的引用，你就需要遍历整个链表找到尾部节点，时间复杂度是 O(n)；
// 如果有尾部节点的引用，就可以在O(1)的时间复杂度内完成尾部添加元素的操作。

// 关键点2: 虚拟头尾节点
// 就是在创建双链表时就创建一个虚拟头节点和一个虚拟尾节点，无论双链表是否为空，这两个节点都存在。这样就不会出现空指针的问题，可以避免很多边界情况的处理
// 有虚拟头尾节点的空链表：dummyHead <-> dummyTail
// 以前要把在头部插入元素、在尾部插入元素和在中间插入元素几种情况分开讨论，现在有了头尾虚拟节点，无论链表是否为空，都只需要考虑在中间插入元素的情况就可以了，这样代码会简洁很多。
// 当然，虚拟头结点会多占用一点内存空间，但是比起给你解决的麻烦，这点空间消耗是划算的

// 关键点3: 内存泄漏
// 垃圾回收的判断机制是看这个对象是否被别人引用，而并不会 care 这个对象是否还引用着别人, 所以即使不置为nil, 该节点只要没有引用了，最后也会被垃圾回收机制给回收掉
// 不过出于习惯，最好还是把被删除的节点置为nil,避免一些潜在的问题

// leetCode 707: https://leetcode.cn/problems/design-linked-list/ 设计链表

// Node 单链表节点
type Node[E any] struct {
	val  E
	next *Node[E]
}

// SinglyLinkedList 单链表实现
type SinglyLinkedList[E any] struct {
	head *Node[E]
	tail *Node[E]
	size int
}

func NewSinglyLinkedList[E any]() *SinglyLinkedList[E] {
	head := &Node[E]{}
	return &SinglyLinkedList[E]{head: head, tail: head, size: 0}
}

// AddFirst 在链表头部添加元素
func (s *SinglyLinkedList[E]) AddFirst(e E) {
	newNode := &Node[E]{val: e}
	newNode.next = s.head.next
	s.head.next = newNode
	if s.size == 0 {
		s.tail = newNode
	}
	s.size++
}

// AddLast 在链表尾部添加元素
func (s *SinglyLinkedList[E]) AddLast(e E) {
	newNode := &Node[E]{val: e}
	s.tail.next = newNode
	s.tail = newNode
	s.size++
}

// Add 在指定位置添加元素
func (s *SinglyLinkedList[E]) Add(index int, element E) error {
	if index < 0 || index > s.size {
		return errors.New("index out of bounds")
	}

	if index == s.size {
		s.AddLast(element)
		return nil
	}

	prev := s.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	newNode := &Node[E]{val: element}
	newNode.next = prev.next
	prev.next = newNode
	s.size++
	return nil
}

// RemoveFirst 移除链表头部元素
func (s *SinglyLinkedList[E]) RemoveFirst() (E, error) {
	if s.IsEmpty() {
		return *new(E), errors.New("no such element")
	}
	first := s.head.next
	s.head.next = first.next
	s.size--
	if s.size == 0 {
		s.tail = s.head
	}
	return first.val, nil
}

// RemoveLast 移除链表尾部元素
func (s *SinglyLinkedList[E]) RemoveLast() (E, error) {
	if s.IsEmpty() {
		return *new(E), errors.New("no such element")
	}

	prev := s.head
	for prev.next != s.tail {
		prev = prev.next
	}
	val := s.tail.val
	prev.next = nil
	s.tail = prev
	s.size--
	return val, nil
}

// Remove 移除指定位置的元素
func (s *SinglyLinkedList[E]) Remove(index int) (E, error) {
	if index < 0 || index >= s.size {
		return *new(E), errors.New("index out of bounds")
	}

	prev := s.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	nodeToRemove := prev.next
	prev.next = nodeToRemove.next
	if index == s.size-1 {
		s.tail = prev
	}
	s.size--
	return nodeToRemove.val, nil
}

// GetFirst 获取链表头部元素
func (s *SinglyLinkedList[E]) GetFirst() (E, error) {
	if s.IsEmpty() {
		return *new(E), errors.New("no such element")
	}
	return s.head.next.val, nil
}

// GetLast 获取链表尾部元素
func (s *SinglyLinkedList[E]) GetLast() (E, error) {
	if s.IsEmpty() {
		return *new(E), errors.New("no such element")
	}
	return s.Get(s.size - 1)
}

// Get 获取指定位置的元素
func (s *SinglyLinkedList[E]) Get(index int) (E, error) {
	if index < 0 || index >= s.size {
		return *new(E), errors.New("index out of bounds")
	}
	node := s.getNode(index)
	return node.val, nil
}

// Set 设置指定位置的元素
func (s *SinglyLinkedList[E]) Set(index int, element E) (E, error) {
	if index < 0 || index >= s.size {
		return *new(E), errors.New("index out of bounds")
	}
	node := s.getNode(index)
	oldVal := node.val
	node.val = element
	return oldVal, nil
}

// Size 获取链表大小
func (s *SinglyLinkedList[E]) Size() int {
	return s.size
}

// IsEmpty 检查链表是否为空
func (s *SinglyLinkedList[E]) IsEmpty() bool {
	return s.size == 0
}

// getNode 获取指定索引的节点
func (s *SinglyLinkedList[E]) getNode(index int) *Node[E] {
	p := s.head.next
	for i := 0; i < index; i++ {
		p = p.next
	}
	return p
}

// Print 打印链表中的所有值
func (s *SinglyLinkedList[E]) Print() {
	current := s.head
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}
