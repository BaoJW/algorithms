package list

import (
	"container/heap"
)

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

// 链表类题目3：合并k个有序链表
// https://leetcode.cn/problems/merge-k-sorted-lists/description/
// 给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。

// 示例 1：
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
// 	1->4->5,
// 	1->3->4,
// 	2->6
// ]
// 将它们合并到一个有序链表中得到。
//	1->1->2->3->4->4->5->6

// 示例 2：
//	输入：lists = []
//	输出：[]

// 示例 3：
//	输入：lists = [[]]
//	输出：[]

// 核心思想：我们如何快速得到k个节点中的最小节点，接到结果链表上
// 自然而然可以想到，寻找k个节点中最小节点，最好的方式是排序好之后取最小值，我并不需要k个节点完全的顺序，我只需要找到里面的最小值即可
// 自然而然想到小顶堆，完全贴合优先级队列的实现
// ***面试常考题
// 时间复杂度：优先队列pq中的元素个数最多是k，所以一次pop或者push方法的时间复杂度是O(logK),所有链表节点都会被加入和弹出pq，所以算法整体的时间复杂度是O(NlogK)，其中k是链表条数，N是这些链表的节点总数

// 为Node实现heap.Interface接口

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

// MergeKLists 合并 k 个升序链表
func MergeKLists(lists []*Node) *Node {
	if len(lists) == 0 {
		return nil
	}

	// 虚拟头结点
	dummy := &Node{-1, nil}
	p := dummy

	// 优先级队列，最小堆
	pq := &PriorityQueue{}
	heap.Init(pq)

	// 将 k 个链表的头结点加入最小堆
	for _, head := range lists {
		if head != nil {
			heap.Push(pq, head)
		}
	}

	for pq.Len() > 0 {
		// 获取最小节点，接到结果链表中
		node := heap.Pop(pq).(*Node)
		p.Next = node
		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
		// p 指针不断前进
		p = p.Next
	}
	return dummy.Next
}
