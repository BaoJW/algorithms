package tree

import "fmt"

// 最常见的二叉树就是类似链表那样的链式存储结构，每个二叉树节点有指向左右子节点的指针，这种方式比较简单直观

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// State 定义状态结构
type State struct {
	node  *TreeNode
	depth int
}

func Constrcut() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}
}

// 构建出来的二叉树是这样的：
//     1
//    / \
//   2   3
//  /   / \
// 4   5   6

// 二叉树的递归/层序遍历
// 二叉树的遍历分成三种，按照根节点的访问先后分为:
// 前序遍历(先根遍历): 先访问根节点，然后访问左子树，最后访问右子树
// 中序遍历(中根遍历): 先访问左子树，然后访问根节点，最后访问右子树
// 后序遍历(后根遍历): 先访问左子树，然后访问右子树，最后访问根节点

// 二叉树的遍历框架
//func traverse(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	// 前序位置
//	traverse(root.Left)
//	// 中序位置
//	traverse(root.Right)
//	// 后序位置
//}

// 创建树节点
func createTree(nodes []interface{}) *TreeNode {
	if len(nodes) == 0 || nodes[0] == nil {
		return nil
	}

	root := &TreeNode{Val: nodes[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(nodes) {
		current := queue[0]
		queue = queue[1:]

		if nodes[i] != nil {
			current.Left = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, current.Left)
		}
		i++

		if i < len(nodes) && nodes[i] != nil {
			current.Right = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, current.Right)
		}
		i++
	}
	return root
}

// 遍历函数 递归遍历(DFS)
func traverse(root *TreeNode, preorderResult, inorderResult, postorderResult *[]int) {
	if root == nil {
		return
	}
	// 前序遍历
	*preorderResult = append(*preorderResult, root.Val)
	// 递归遍历左子树
	traverse(root.Left, preorderResult, inorderResult, postorderResult)
	// 中序遍历
	*inorderResult = append(*inorderResult, root.Val)
	// 递归遍历右子树
	traverse(root.Right, preorderResult, inorderResult, postorderResult)
	// 后序遍历
	*postorderResult = append(*postorderResult, root.Val)
}

// 遍历函数 层序遍历(BFS)
func levelOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	q := []State{{node: root, depth: 1}}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		// 访问 cur 节点，同时知道它的路径权重和
		fmt.Printf("depth = %d, val = %d\n", cur.depth, cur.node.Val)

		// 把 cur 的左右子节点加入队列
		if cur.node.Left != nil {
			q = append(q, State{node: cur.node.Left, depth: cur.depth + 1})
		}
		if cur.node.Right != nil {
			q = append(q, State{node: cur.node.Right, depth: cur.depth + 1})
		}
	}
}
