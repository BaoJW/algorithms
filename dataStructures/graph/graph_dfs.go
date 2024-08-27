package graph

import "fmt"

// 图结构深度优先搜索(DFS)
// 遍历所有节点(visited数组)
func traverse(graph Graph, s int, visited []bool) {
	// base case
	if s < 0 || s >= graph.Size() {
		return
	}
	if visited[s] {
		// 防止死循环
		return
	}
	// 前序位置
	visited[s] = true
	for _, e := range graph.Neighbors(s) {
		traverse(graph, e.to, visited)
	}
	// 后序位置

}

// 由于visited数组的剪枝作用，这个遍历函数会遍历一次图中的所有节点，并尝试遍历一次所有边，所以算法的时间复杂度是O(E+V),其中E是边的总数，V的节点的总数
// 时间复杂度为什么是O(E+V)？
// 二叉树/多叉树的遍历函数，也要算上边的数量，只不过对于树结构来说，边的数量和节点的数量是近似相等的，所以时间复杂度还是O(N+N) = O(N)
// 树结构中的边只能由父节点指向子节点，所以除了根节点，你可以把每个节点和它上面那条来自父节点的边配成一对儿，这样就可以比较直观地看出边的数量和节点的数量是近似相等的
// 而对于图结构来说，任意两个节点之间都可以连接一条边，所以边的数量和节点的数量不再有特定的关系，所以我们要说图的遍历函数时间复杂度是O(E+V)

// 遍历所有路径(onPath数组)
// 对于树结构，遍历所有路径和遍历所有节点是没什么区别的，因为对于树结构来说，只能由父节点指向子节点，所以从根节点root出发，到任意一个节点targetNode的路径是唯一的；换句话说，我遍历一遍树结构的所有节点之后，必然可以找到root到targetNode的唯一路径
// 对于图结构，由起点src到目标节点dest的路径可能不止一条，我们需要一个onPath数组，在进入节点时(前序位置)标记为正在访问,退出节点时(后续位置)撤销标记,这样才能遍历图中的所有路径,从而找到src到dest的所有路径
const size = 10

var onPath []bool = make([]bool, size)
var path []int

func traverseOnPath(graph Graph, src, dest int) {
	if src < 0 || src >= size {
		return
	}

	if onPath[src] {
		// 防止死循环(成环)
		return
	}

	// 前序位置
	onPath[src] = true
	path = append(path, src)
	if src == dest {
		fmt.Println("find path:", path)
	}
	for _, e := range graph.Neighbors(src) {
		traverseOnPath(graph, e.to, dest)
	}

	// 后序位置
	path = path[:len(path)-1]
	onPath[src] = false
}

// 同时使用visited和onPath数组
// 由于遍历所有路径的算法复杂度较高，某些路径遍历的场景下我们可能会借助visited数组进行剪枝，提前排除一些不符合条件的路径，从而降低复杂度
// 比如拓扑排序中会讲到如何判定图是否成环，就会同时利用visited和onPath数组来进行剪枝
// 比如说判定成环的场景，在遍历所有路径的过程中，如果一个节点s被标记为visited，那么说明从s这个起点出发的所有路径在之前都已经遍历过了，如果之前遍历的时候都没有找到环，我现在再去遍历一次，肯定也会找不到，所以这里就可以直接剪枝，不再继续遍历节点s
