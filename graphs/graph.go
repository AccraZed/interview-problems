package graphs

type Node struct {
	Val       int
	Neighbors []*Node
}

// newGraph creates a node-based adjacency list representation of nodes, given
// an adjacency matrix representation of the graph. For each index i, graph[i]
// is a list of all nodes j that connect to node i
func newGraph(graph [][]int) []*Node {
	nodes := make([]*Node, len(graph))
	for i := range nodes {
		nodes[i] = new(Node)
		nodes[i].Val = i
	}

	for i, node := range graph {
		for _, neighbor := range node {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neighbor])
		}
	}

	return nodes
}
