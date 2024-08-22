package tree

import "fmt"

// 多叉树

type Node struct {
	Val      int
	Children []*Node
}

// 多叉树的递归遍历(DFS)
// 多叉树没有中序位置，因为可能有多个节点，所谓的中序位置也就没什么意义了

// 创建 N 叉树节点
func createNode(nodes []interface{}) *Node {
	if len(nodes) == 0 || nodes[0] == nil {
		return nil
	}

	root := &Node{Val: nodes[0].(int)}
	queue := []*Node{root}
	i := 1

	for len(queue) > 0 && i < len(nodes) {
		// 跳过 nil
		if nodes[i] == nil {
			i++
			continue
		}
		current := queue[0]
		queue = queue[1:]

		// 处理当前节点的子节点
		for i < len(nodes) && nodes[i] != nil {
			child := &Node{Val: nodes[i].(int)}
			current.Children = append(current.Children, child)
			queue = append(queue, child)
			i++
		}
	}
	return root
}

// 前序和后序遍历结果切片
var preorderResult []int
var postorderResult []int

// 遍历函数 多叉树的递归遍历
func traverseN(root *Node) {
	if root == nil {
		return
	}
	// 前序遍历位置
	preorderResult = append(preorderResult, root.Val)
	for _, child := range root.Children {
		traverseN(child)
	}
	// 后序遍历位置
	postorderResult = append(postorderResult, root.Val)
}

// 遍历函数 多叉树的层序遍历
type StateN struct {
	node  *Node
	depth int
}

func levelOrderTraverseN(root *Node) {
	if root == nil {
		return
	}
	q := []StateN{}
	// 记录当前遍历到的层数（根节点视为第 1 层）
	q = append(q, StateN{root, 1})

	for len(q) > 0 {
		state := q[0]
		q = q[1:]
		cur := state.node
		depth := state.depth
		// 访问 cur 节点，同时知道它所在的层数
		fmt.Printf("depth = %d, val = %d\n", depth, cur.Val)

		for _, child := range cur.Children {
			q = append(q, StateN{child, depth + 1})
		}
	}
}
