package list

// 链表类题目1：合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/
// 将两个升序链表合并为一个新的升序链表并返回，新链表是通过拼接给定的两个链表的所有节点组成

// 示例1:
// 输入: l1 = [1,2,4], l2 = [1,3,4]
// 输出: [1,1,2,3,4,4]

// 示例2:
// 输入: l1 = [], l2 = []
// 输出: []

// 示例3:
// 输入: l1 = [], l2 = [0]
// 输出: [0]

type Node struct {
	Val  int
	Next *Node
}

type SingleList struct {
	head *Node
	tail *Node
	size int
}

func NewSingleList() *SingleList {
	head := &Node{}
	return &SingleList{head: head, tail: head, size: 0}
}

func (s *SingleList) AddElementToLast(val int) {
	newNode := &Node{Val: val}
	s.tail.Next = newNode
	s.tail = newNode
	s.size++
}

// getNode 获取指定索引的节点
func (s *SingleList) getNode(index int) *Node {
	p := s.head.Next
	for i := 0; i < index; i++ {
		p = p.Next
	}
	return p
}

// MergeList 这个是自己没看答案写出来的，但是有些问题没有思考到位
// 1. newList相当于是新建了一条链表，这对内存是损耗，其实对于本题来说没有必要
// 2. 本质上操作原始链表的指针就能完成
func MergeList(listA, listB *SingleList) (newList *SingleList) {
	// 第一种情况，两个链表都为空，那就返回一个空链表回去
	if listA.size == 0 && listB.size == 0 {
		return NewSingleList()
	}

	// 第二种情况，一个链表为空，一个链表不为空
	if listA.size == 0 && listB.size != 0 {
		return listB
	}

	// 第二种情况，一个链表为空，一个链表不为空
	if listB.size == 0 && listA.size != 0 {
		return listA
	}

	// 第三种情况，两个链表都不为空
	newList = NewSingleList()
	pointA := listA.head.Next
	pointB := listB.head.Next
	pointNew := newList.head
	totalSize := listA.size + listB.size

	for i := 1; i <= totalSize; i++ {
		switch {
		case pointA == nil && pointB != nil:
			pointNew.Next = pointB
			pointB = pointB.Next
		case pointB == nil && pointA != nil:
			pointNew.Next = pointA
			pointA = pointA.Next
		case pointA.Val <= pointB.Val:
			pointNew.Next = pointA
			pointA = pointA.Next
		default:
			pointNew.Next = pointB
			pointB = pointB.Next
		}
		pointNew = pointNew.Next

	}

	return newList
}

// 辅助函数：创建链表
func createList(elements []int) *Node {
	dummy := &Node{}
	current := dummy
	for _, val := range elements {
		current.Next = &Node{Val: val}
		current = current.Next
	}
	return dummy.Next
}

// 辅助函数：提取链表中的值
func extractValues(head *Node) []int {
	var result []int
	current := head
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

// MergeListV2 基于答案去优化的方法
func MergeListV2(listA, listB *Node) *Node {
	dummy := &Node{} // 作为虚拟头节点
	pointNew := dummy
	pointA := listA
	pointB := listB
	for pointA != nil && pointB != nil {
		// 比较 pointA和pointB两个指针
		// 将值较小的的节点接到point指针
		if pointA.Val <= pointB.Val {
			pointNew.Next = pointA
			pointA = pointA.Next
		} else {
			pointNew.Next = pointB
			pointB = pointB.Next
		}
		// pointNew指针不断前进
		pointNew = pointNew.Next
	}

	if pointA != nil {
		pointNew.Next = pointA
	}

	if pointB != nil {
		pointNew.Next = pointB
	}

	return dummy.Next

}
