package graph

// 图：多叉树的延升
// 在树结构中，只允许父节点指向子节点,不存在子节点指向父节点的情况，子节点之间也不会互相连接
// 在图结构中，节点之间可以互指，形成复杂的网络结构

// 一幅图是由节点(Vertex)和边(Edge)组成的
// "度"的概念：
// 在无向图中，度就是每个节点相连的边的条数
// 由于有向图中每个节点的度被细分为入度(indegree)和出度(outdegree)

// 图结构的两种实现方法：邻接表和邻接矩阵
// 邻接表：就是把每个节点x的邻居都存到一个列表里，然后把x和这个列表映射起来，这样就可以通过一个节点x找到他的所有相邻节点
// 具体实现：var graph [][]int  graph[x] 存储 x 的所有邻居节点
// 邻接矩阵：一个二维布尔数组，权且称为matrix,如果节点x和节点y是相连的，则matrix[x][y] = true,如果想找x节点的邻居，则遍历matrix[x]即可
// 具体实现：matrix[x][y] 记录 x 是否有一条指向 y 的边

// 邻接表和邻接矩阵的使用场景：
// 注意分析两种存储方式的空间复杂度，对于一个有V个节点，E条边的图，邻接表的空间复杂度是O(V+E), 而邻接矩阵的空间复杂度是O(V^2)
// 所以如果一幅图的 E 远小于 V^2（稀疏图），那么邻接表会比邻接矩阵节省空间，反之，如果 E 接近 V^2（稠密图），二者就差不多了
// 邻接矩阵的最大优势在于，矩阵是一个强有力的数学工具，图的一些隐晦性质可以借助精妙的矩阵运算展现出来

// 图结构的DFS-BFS遍历
// 一句话总结: 图的遍历就是多叉树遍历的延伸，主要的遍历方式还是深度优先搜索(DFS)和广度优先搜索(BFS)
// 唯一的区别是：树结构中不存在环，而图结构中可能存在环，所以我们需要标记遍历过的节点，避免遍历函数在环中死循环
// 具体实现1：遍历图的所有节点时，需要visited数组在前序位置标记节点
// 具体实现2: 遍历图的所有路径时，需要onPath数组在前序位置标记节点，在后续未知撤销标记

// Edge 存储相邻节点及边的权重
type Edge struct {
	to     int
	weight int
}

type Graph interface {
	AddEdge(from, in, weight int) // 添加一条边(带权重)
	RemoveEdge(from, to int)      // 删除一条边
	HasEdge(from, to int) bool    // 判断两个节点是否相连
	Weight(from, to int) int      // 返回一条边的权重
	Neighbors(v int) []Edge       // 返回某个节点的所有邻居节点和对应权重
	Size() int                    // 返回节点总数
}

// WeightedDigraph 有向加权图(邻接表实现)
type WeightedDigraph struct {
	graph [][]Edge
}

func NewWeightedDigraph(n int) *WeightedDigraph {
	// 我们这里简单起见，建图时要传入节点总数，这其实可以优化
	// 比如把 graph 设置为 map[int][]Edge，就可以动态添加新节点了
	graph := make([][]Edge, n)
	return &WeightedDigraph{graph: graph}
}

// AddEdge 增，添加一条带权重的有向边，复杂度 O(1)
func (g *WeightedDigraph) AddEdge(from, to, weight int) {
	g.graph[from] = append(g.graph[from], Edge{to: to, weight: weight})
}

// RemoveEdge 删，删除一条有向边，复杂度 O(V)
func (g *WeightedDigraph) RemoveEdge(from, to int) {
	for i, e := range g.graph[from] {
		if e.to == to {
			g.graph[from] = append(g.graph[from][:i], g.graph[from][i+1:]...)
			break
		}
	}
}

// HasEdge 查，判断两个节点是否相邻，复杂度 O(V)
func (g *WeightedDigraph) HasEdge(from, to int) bool {
	for _, e := range g.graph[from] {
		if e.to == to {
			return true
		}
	}
	return false
}

// Weight 查，返回一条边的权重，复杂度 O(V)
func (g *WeightedDigraph) Weight(from, to int) int {
	for _, e := range g.graph[from] {
		if e.to == to {
			return e.weight
		}
	}
	panic("No such edge")
}

// 上面的 HasEdge、RemoveEdge、Weight 方法遍历 List 的行为是可以优化的
// 比如用 map[int]map[int]int 存储邻接表
// 这样就可以避免遍历 List，复杂度就能降到 O(1)

// Neighbors 查，返回某个节点的所有邻居节点，复杂度 O(1)
func (g *WeightedDigraph) Neighbors(v int) []Edge {
	return g.graph[v]
}

func (g *WeightedDigraph) Size() int {
	return len(g.graph)
}

// WeightedDigraphOfMatrix 有向加权图(邻接矩阵实现)
type WeightedDigraphOfMatrix struct {
	// 邻接矩阵，matrix[from][to] 存储从节点 from 到节点 to 的边的权重
	// 0 表示没有连接
	matrix [][]int
}

func NewWeightedDigraphOfMatrix(n int) *WeightedDigraphOfMatrix {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	return &WeightedDigraphOfMatrix{matrix}
}

// AddEdge 增，添加一条带权重的有向边，复杂度 O(1)
func (g *WeightedDigraphOfMatrix) AddEdge(from, to, weight int) {
	g.matrix[from][to] = weight
}

// RemoveEdge 删，删除一条有向边，复杂度 O(1)
func (g *WeightedDigraphOfMatrix) RemoveEdge(from, to int) {
	g.matrix[from][to] = 0
}

// HasEdge 查，判断两个节点是否相邻，复杂度 O(1)
func (g *WeightedDigraphOfMatrix) HasEdge(from, to int) bool {
	return g.matrix[from][to] != 0
}

// Weight 查，返回一条边的权重，复杂度 O(1)
func (g *WeightedDigraphOfMatrix) Weight(from, to int) int {
	return g.matrix[from][to]
}

// Neighbors 查，返回某个节点的所有邻居节点，复杂度 O(V)
func (g *WeightedDigraphOfMatrix) Neighbors(v int) []Edge {
	res := []Edge{}
	for i := 0; i < len(g.matrix[v]); i++ {
		if g.matrix[v][i] > 0 {
			res = append(res, Edge{to: i, weight: g.matrix[v][i]})
		}
	}
	return res
}

func (g *WeightedDigraphOfMatrix) Size() int {
	return len(g.matrix)
}
