package binaryHeap

// 二叉堆: 一种能够动态排序的数据结构
// 主要操作就两个，sink(下沉)和swim(上浮),用以维护二叉堆的性质
// 主要应用有两个：
// 1. 一种数据结构: 优先级队列(Priority Queue)
// 2. 一种排序方法: 堆排序(Heap Sort)

// 所谓动态排序，就是我们可以不断的往数据结构里添加或删除元素，数据结构会自动调整元素的位置，使得我们可以有序的从数据结构中读取元素，这是类似快速排序/归并排序等排序算法做不到的
// 能够动态排序的常用数据结构其实只有两个，一个是优先级队列(底层用二叉堆实现)，另一个是二叉搜索树，优先级队列能做的事二叉搜索树都可以做，但因为优先级队列的实现更为简单，所以能用优先级队列解决的问题就无需二叉搜索树了

// 二叉堆的性质：
// 二叉堆是一种特殊的二叉树，这棵二叉树上的任意节点的值，都必须大于等于(或小于等于)其左右子树所有节点的值
// 如果是大于等于，我们称之为大顶堆，根节点就是整棵树的最大值
// 如果是小于等于，我们称之为小顶堆，根节点就是整棵树的最小值

// 增：push方法插入元素
// 核心步骤：以小顶堆为例，向小顶堆中插入新元素遵循两个步骤：
// 1. 先把新元素追加到二叉树底层的最右侧，保持完全二叉树的结构，此时该元素的父节点可能比它大，不满足小顶堆的性质
// 2. 为了恢复小顶堆的性质，需要讲这个新元素不断上浮(swim)，直到它的父节点比它小为止，或者到达根节点，此时整个二叉树就满足小顶堆的性质了

// 删: pop方法删除元素
// 核心步骤：以小顶堆为例，删除小顶堆的堆顶元素遵循两个步骤：
// 1. 先把堆顶元素删除，把二叉树底层的最右侧元素摘除并移动到堆顶，保持完全二叉树的结构，此时堆顶元素可能比它的子节点大，不满足小顶堆的性质
// 2. 为了恢复小顶堆的性质，需要将这个新的堆顶元素不断下沉(sink)，直到它比它的子节点小为止，或者到达叶子节点，此时整个二叉树就满足小顶堆的性质了

// 查: peek方法查看堆顶元素

// 在数组上模拟二叉堆
// 原因:
// 1. 链表节点需要一个额外的指针存储相邻节点的地址，所以相对数组，链表的内存消耗会大一些
// 2. 时间复杂度，如果使用链表，想要拿到二叉堆最底层的最右侧元素，那么需要DFS/BFS二叉树，时间复杂度是O(N)，进而导致pop/push方法的时间复杂度都退化到O(N)
// 用数组模拟二叉堆，前提是这个二叉树必须要是完全二叉树(除了最后一层，其他层的节点都是满的，最后一层的节点都靠左排列)

// MyPriorityQueue 定义优先队列结构体
type MyPriorityQueue struct {
	heap []int
	size int
	cmp  func(int, int) bool // 比较函数
}

// NewMyPriorityQueue 创建新的优先队列
func NewMyPriorityQueue(capacity int, cmp func(int, int) bool) *MyPriorityQueue {
	return &MyPriorityQueue{
		heap: make([]int, capacity+1), // 多分配一个空间，索引0不使用
		size: 0,
		cmp:  cmp,
	}
}

// 获取父节点的索引
func (pq *MyPriorityQueue) parent(node int) int {
	return node / 2
}

// 获取左子节点的索引
func (pq *MyPriorityQueue) left(node int) int {
	return node * 2
}

// 获取右子节点的索引
func (pq *MyPriorityQueue) right(node int) int {
	return node*2 + 1
}

// 交换堆中两个元素
func (pq *MyPriorityQueue) swap(i, j int) {
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
}

// Peek 返回堆顶元素，时间复杂度 O(1)
func (pq *MyPriorityQueue) Peek() int {
	if pq.size == 0 {
		panic("Priority queue underflow")
	}
	return pq.heap[1]
}

// Push 向堆中插入一个元素，时间复杂度 O(logN)
func (pq *MyPriorityQueue) Push(x int) {
	// 扩容
	if pq.size == len(pq.heap)-1 {
		pq.resize(2 * len(pq.heap))
	}
	pq.size++
	pq.heap[pq.size] = x
	pq.swim(pq.size) // 上浮到正确位置
}

// Pop 删除堆顶元素，时间复杂度 O(logN)
func (pq *MyPriorityQueue) Pop() int {
	if pq.size == 0 {
		panic("Priority queue underflow")
	}
	res := pq.heap[1]
	pq.swap(1, pq.size) // 将堆底元素放到堆顶
	pq.size--
	pq.sink(1)             // 下沉到正确位置
	pq.heap[pq.size+1] = 0 // 避免对象游离
	if pq.size > 0 && pq.size == (len(pq.heap)-1)/4 {
		pq.resize(len(pq.heap) / 2) // 缩容
	}
	return res
}

// 上浮操作，时间复杂度 O(logN)
func (pq *MyPriorityQueue) swim(k int) {
	for k > 1 && pq.cmp(pq.heap[pq.parent(k)], pq.heap[k]) {
		pq.swap(pq.parent(k), k)
		k = pq.parent(k)
	}
}

// 下沉操作，时间复杂度 O(logN)
func (pq *MyPriorityQueue) sink(k int) {
	for pq.left(k) <= pq.size {
		j := pq.left(k)
		if j < pq.size && pq.cmp(pq.heap[j], pq.heap[j+1]) {
			j++
		}
		if !pq.cmp(pq.heap[k], pq.heap[j]) {
			break
		}
		pq.swap(k, j)
		k = j
	}
}

// 扩容方法
func (pq *MyPriorityQueue) resize(capacity int) {
	newHeap := make([]int, capacity)
	copy(newHeap, pq.heap)
	pq.heap = newHeap
}
