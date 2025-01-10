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

/***
第二题：环形链表
给你一个链表的头节点 head ，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
如果链表中存在环 ，则返回 true 。 否则，返回 false 。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。

提示：
链表中节点的数目范围是 [0, 104]
-105 <= Node.val <= 105
pos 为 -1 或者链表中的一个 有效索引 。

进阶：你能用 O(1)（即，常量）内存解决此问题吗？
***/

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}
