package twoPoints

type ListNode struct {
	Val  int
	Next *ListNode
}

// 本文的算法题主要是使用双指针解决链表的问题技巧

/***
第一题：链表的中间结点
	给你单链表的头结点 head ，请你找出并返回链表的中间结点。
	如果有两个中间结点，则返回第二个中间结点。
示例 1：
输入：head = [1,2,3,4,5]
输出：[3,4,5]
解释：链表只有一个中间结点，值为 3 。

示例 2：
输入：head = [1,2,3,4,5,6]
输出：[4,5,6]
解释：该链表有两个中间结点，值分别为 3 和 4 ，返回第二个结点。

提示：
链表的结点数范围是 [1, 100]
1 <= Node.val <= 100
***/

/*
	解题思路：利用双指针一个的速度是另一个的一半，当fast到达链表尾部的时候，slow就到了链表的中间，此时slow就是我们想要的中间节点
*/

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
