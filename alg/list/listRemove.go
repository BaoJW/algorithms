package list

// 链表类题目4：删除链表的倒数第N个节点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
// 给你一个链表，删除链表的倒数第N个节点，并返回链表的头节点

// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
// 示例 2：

// 输入：head = [1], n = 1
// 输出：[]

// 示例 3：
// 输入：head = [1,2], n = 1
// 输出：[1]

// 核心思想: 单链表的倒数第k个节点
// 暴力模式：两次循环实现，第一次遍历链表获得链表的总长度(假设为N个节点，长度为N),那么倒数第k个节点就是正数的(N-k+1)个节点，所以第二次遍历到(N-k+1)个节点即为所求，从时间复杂度上来看，如果用bigO表示法来计算时间复杂度，无论1次遍历还是2次遍历他的时间复杂度都是O(N)
// 技巧模式：
// 1. 先让一个指针pointA指向链表的头节点head,然后走k步，现在的pointA只要在走n-k步，就能走到链表末尾的空指针
// 2. 在pointA走了k步后，再让一个指针pointB指向链表的头节点head, 然后让pointA和pointB同时向前走
// 3. pointA走到链表末尾的空指针时前进了n-k步，pointB也从head开始前进了n-k步，停留在第n-k+1个节点上，即恰好停在了链表倒数第k个节点上
func removeNthFromEnd(head *Node, n int) *Node {
	// 虚拟头节点
	dummy := &Node{Val: -1}
	dummy.Next = head
	// 删除倒数第n个，要先找到倒数第n+1个节点
	findNode := findFromEnd(dummy, n+1)
	// 删除倒数第n个节点
	findNode.Next = findNode.Next.Next
	return dummy.Next
}

// 返回链表的倒数第 k 个节点
func findFromEnd(head *Node, k int) *Node {
	pointA := head
	// p1先走k步
	for i := 0; i < k; i++ {
		pointA = pointA.Next
	}

	pointB := head
	// p1和p2同时走n-k步
	for pointA != nil {
		pointA = pointA.Next
		pointB = pointB.Next
	}

	// p2指向第n-k+1个节点，即倒数第k个节点
	return pointB
}

// 链表类题目5：链表的中间节点
// https://leetcode.cn/problems/middle-of-the-linked-list/description/
// 给你单链表的头节点 head ，请你找出并返回链表的中间节点
// 如果有两个中间节点，则返回第二个中间节点

// 示例1:
// 输入：head = [1,2,3,4,5]
// 输出：[3,4,5]
// 解释：链表只有一个中间节点，值为3

// 示例2:
// 输入：head = [1,2,3,4,5,6]
// 输出：[4,5,6]
// 解释：该链表有两个中间节点,值分别为3和4,返回第二个节点

// 核心思想: 快慢指针技巧
// 1. 我们让两个指针slow和fast分别指向链表头节点head
// 2. 每当慢指针slow前进一步，快指针fast就前进两步，这样当fast走到链表末尾时，slow就指向了链表重点

func middleNode(head *Node) *Node {
	// 快慢指针初始化指向 head
	pointFast, pointSlow := head, head
	// 快指针走到末尾时停止
	for pointFast != nil && pointFast.Next != nil {
		// 慢指针走一步，快指针走两步
		pointFast = pointFast.Next.Next
		pointSlow = pointSlow.Next
	}

	// 慢指针指向中点
	return pointSlow
}

// 判断链表是否包含环
// 核心思想：快慢指针技巧
// 每当慢指针slow前进一步，快指针fast就前进两步
// 如果fast最终能正常走到链表末尾，说明链表中没有环；如果fast走着走着竟然和slow相遇了，那肯定是fast在链表中转圈了，说明链表中含有环

func hasCycle(head *Node) bool {
	// 快慢指针初始化指向 head
	pointFast, pointSlow := head, head
	// 快指针走到末尾时停止
	for pointFast != nil && pointFast.Next != nil {
		// 慢指针走一步，快指针走两步
		pointFast = pointFast.Next.Next
		pointSlow = pointSlow.Next

		// 快慢指针相遇，说明含有环
		if pointFast == pointSlow {
			return true
		}
	}

	return false
}

// 链表类题目6：环形链表
// https://leetcode.cn/problems/linked-list-cycle-ii/
// 给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
// 不允许修改链表。

// 示例1
// 输入：head = [3,2,0,-4], pos = 1
// 输出：返回索引为 1 的链表节点
// 解释：链表中有一个环，其尾部连接到第二个节点。

// 示例2
// 输入：head = [1,2], pos = 0
// 输出：返回索引为 0 的链表节点
// 解释：链表中有一个环，其尾部连接到第一个节点

// 示例3
// 输入：head = [1], pos = -1
// 输出：返回 null
// 解释：链表中没有环。

// 核心思想：快慢指针
// 假设快慢指针相遇时，慢指针slow走了k步，那么快指针fast一定走了2k步
// fast一定比slow多走了k不，这多走的k步其实就是fast指针在环里转圈圈，所以k得值就是环长度的整数倍
// 假设相遇点距环的起点的距离为m，那么结合上图的slow指针，环的起点距头节点head的距离为k-m,也就是说如果从head前进k-m步就能到达环起点
// 巧的是，如果从相遇点继续前进k-m步，也恰好到达环起点，因为结合上图的fast指针，从相遇点开始走k步可以转回到相遇点，那走k-m步肯定就走到环起点了
// 所以，只要我们把快慢指针中的任一个重新指向head,然后两个指针同速前进，k-m 步后一定会相遇，相遇之处就是环的起点了

func detectCycle(head *Node) *Node {
	// 快慢指针初始化指向 head
	pointFast, pointSlow := head, head
	// 快指针走到末尾时停止
	for pointFast != nil && pointFast.Next != nil {
		// 慢指针走一步，快指针走两步
		pointFast = pointFast.Next.Next
		pointSlow = pointSlow.Next

		// 快慢指针相遇，说明含有环
		if pointFast == pointSlow {
			pointA := head
			for pointA != pointSlow {
				pointA = pointA.Next
				pointSlow = pointSlow.Next
			}
			return pointA
		}
	}

	return nil

}

// 链表类题目7：相交链表
// https://leetcode.cn/problems/intersection-of-two-linked-lists/
// 给你两个链表的头节点headA和headB，这两个链表可能存在相交
// 如果相交，返回相交的那个节点，如果没相交，返回nil

// 核心思想：通过某些方式，让p1和p2能够同时到达相交节点c1
// 我们可以让p1遍历完链表A之后开始遍历链表B，让p2遍历完链表B之后开始遍历链表A，这样相当于逻辑上两条链表接在了一起
// 如果这样进行拼接，就可以让p1和p2同时进入公共部分，也就是同时到达相交节点c1
func getIntersectionNode(headA, headB *Node) *Node {
	// p1 指向 A 链表头结点，p2 指向 B 链表头结点
	pointA := headA
	pointB := headB
	for pointA != pointB {
		// pointA 走一步，如果走到 A 链表末尾，转到 B 链表
		if pointA == nil {
			pointA = headB
		} else {
			pointA = pointA.Next
		}
		// p2 走一步，如果走到 B 链表末尾，转到 A 链表
		if pointB == nil {
			pointB = headA
		} else {
			pointB = pointB.Next
		}
	}
	return pointA
}
