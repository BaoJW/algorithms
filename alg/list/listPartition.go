package list

// 链表类题目2：分隔链表
// https://leetcode.cn/problems/partition-list/description/
// 给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
// 你应当 保留 两个分区中每个节点的初始相对位置。

// 示例 1：
// 输入：head = [1,4,3,2,5,2], x = 3
// 输出：[1,2,2,4,3,5]
// 示例 2：
// 输入：head = [2,1], x = 2
// 输出：[1,2]

// 提示：
// 链表中节点的数目在范围 [0, 200] 内
// -100 <= Node.val <= 100
// -200 <= x <= 200

// 核心思想：把原链表一分为二
// 把原链表分为两个小链表，一个链表中的元素大小都小于x, 另一个链表中的元素都大于等于x, 最后把两条链表接到一起
//

func Partition(head *Node, x int) *Node {
	// 存放小于 x 的链表的虚拟头结点
	dummyA := &Node{Val: -1, Next: nil}
	// 存放大于等于 x 的链表的虚拟头结点
	dummyB := &Node{Val: -1, Next: nil}
	// p1, p2 指针负责生成结果链表
	pointA, pointB := dummyA, dummyB
	// p 负责遍历原链表，类似合并两个有序链表的逻辑
	// 这里是将一个链表分解成两个链表
	partitionPoint := head
	for partitionPoint != nil {
		if partitionPoint.Val >= x {
			pointB.Next = partitionPoint
			pointB = pointB.Next
		} else {
			pointA.Next = partitionPoint
			pointA = pointA.Next
		}

		// 不能直接让 p 指针前进，
		// p = p.Next
		// 断开原链表中的每个节点的 next 指针
		// 不断开原链表中的链接的话，生成的新链表会成环
		temp := partitionPoint.Next
		partitionPoint.Next = nil
		partitionPoint = temp
	}

	// 链接两个链表
	pointA.Next = dummyB.Next

	return dummyA.Next

}
