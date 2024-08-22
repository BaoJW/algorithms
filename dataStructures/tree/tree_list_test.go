package tree

import (
	"fmt"
	"testing"
)

// 二叉树的递归遍历(DFS)
func Test_Traverse(t *testing.T) {
	preorderResult := []int{}
	inorderResult := []int{}
	postorderResult := []int{}

	allRoot := createTree([]interface{}{1, 2, 3, nil, 4, 5, 6})
	traverse(allRoot, &preorderResult, &inorderResult, &postorderResult)

	fmt.Println("Preorder:", preorderResult)
	fmt.Println("Inorder:", inorderResult)
	fmt.Println("Postorder:", postorderResult)
}

// 二叉树的层序遍历(BFS)
func Test_LevelOrderTraverse(t *testing.T) {
	allRoot := createTree([]interface{}{3, 9, 20, nil, nil, 15, 7})
	levelOrderTraverse(allRoot)
}

// 多叉树的递归遍历(DFS)
func Test_TraverseN(t *testing.T) {
	allRoot := createNode([]interface{}{1, nil, 3, 2, 4, nil, 5, 6})
	traverseN(allRoot)

	fmt.Println("Preorder:", preorderResult)
	fmt.Println("Postorder:", postorderResult)

}
