package graph

import (
	"container/list"
	"fmt"
)

// 图结构的广度优先搜索(BFS)
// 本质就是多叉树的层序遍历，多加一个visited数组来避免重复遍历节点
// BFS算法一般只用来寻找那条最短路径，而不用来求所有路径(这个用DFS比较合适)

// DFS在寻找所有路径时的优势：
// 1. 路径的构建: DFS通过深入每一条路径，直到无法继续为止，这样可以更自然的构建出路径，每当到达一个节点时，可以将当前路径记录下来
// 2. 回溯: DFS方便实现回溯，当到达一个终点或死胡同时，可以轻松返回上一个节点，继续探索其他可能得路径，这种特性使得DFS在处理路径问题时更为高效
// 3. 空间复杂度: 在某些情况下，DFS的空间复杂度可能更低，因为他只需要存储一条路径，而BFS需要存储所有当前层的节点
// 4. 适应性：对于一些特定的图(如树结构)，DFS能够更快的找到所有路径，因为他会优先探索每条边

// BFS在寻找最短路径上的优势：
// 1. 层级遍历: BFS从起始节点开始，逐层探索所有相邻节点，这意味着在访问到某个节点时，所有到达该节点的路径都是通过最少的边数完成的
// 2. 最短路径保证: 由于BFS是按层级进行的，它首先访问距离起始节点最近的所有节点，因此在第一次到达目标节点时，所经过的路径就是最短路径
// 3. 无权图的有效性：BFS特别适合于无权图，因为他不受权重影响，始终能找到最短路径
// 4. 广泛应用: 在许多实际问题中，例如社交网络、地图导航等，BFS能够有效的找到最短路径，尤其是在图的结构较为简单时

// 写法1: 不记录深度
// 图结构的 BFS 遍历，从节点 s 开始进行 BFS
func bfs(graph Graph, s int) {
	visited := make([]bool, graph.Size())
	q := list.New()
	q.PushBack(s)
	visited[s] = true

	for q.Len() > 0 {
		e := q.Front()
		cur := e.Value.(int)
		q.Remove(e)
		fmt.Println("visit", cur)

		for _, edge := range graph.Neighbors(cur) {
			if !visited[edge.to] {
				q.PushBack(edge.to)
				visited[edge.to] = true
			}
		}
	}
}

// 写法2：记录深度
// BFS 遍历图的所有节点，从 s 开始进行 BFS，且记录遍历的步数
func bfsV2(graph Graph, s int) {
	visited := make([]bool, graph.Size())
	q := list.New()
	q.PushBack(s)
	visited[s] = true

	// 记录遍历的步数
	step := 0

	for q.Len() > 0 {
		e := q.Front()
		cur := e.Value.(int)
		q.Remove(e)
		fmt.Printf("visit %d at step %d\n", cur, step)

		for _, edge := range graph.Neighbors(cur) {
			if !visited[edge.to] {
				q.PushBack(edge.to)
				visited[edge.to] = true
			}
		}
		step++
	}
}

// 写法3: 适配不同权重边

// State 定义状态结构
type State struct {
	node   int // 当前节点 ID
	weight int // 从起点 s 到当前节点的权重和
}

// BFS 遍历图的所有节点，从 s 开始进行 BFS，且记录路径的权重和
func bfsV3(graph Graph, s int) {
	visited := make([]bool, graph.Size())
	q := list.New()

	q.PushBack(State{node: s, weight: 0})
	visited[s] = true

	for q.Len() > 0 {
		e := q.Front()
		state := e.Value.(State)
		q.Remove(e)

		cur := state.node
		weight := state.weight
		fmt.Printf("visit %d with path weight %d\n", cur, weight)

		for _, edge := range graph.Neighbors(cur) {
			if !visited[edge.to] {
				q.PushBack(State{node: edge.to, weight: weight + edge.weight})
				visited[edge.to] = true
			}
		}
	}
}
