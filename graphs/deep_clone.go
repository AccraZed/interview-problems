package graphs

import "fmt"

func DeepCloneMain() {
	graph := newGraph([][]int{{1, 2}, {0}, {}})

	clone := clone(graph[0])
	fmt.Println(clone)
}

func clone(node *Node) *Node {
	return deepClone(node, make(map[int]*Node))
}

func deepClone(root *Node, nodes map[int]*Node) *Node {
	// If copy already exists, return it
	if copy, ok := nodes[root.Val]; ok {
		return copy
	}

	copy := new(Node)
	copy.Val = root.Val
	nodes[copy.Val] = copy
	for _, neighbor := range root.Neighbors {
		copy.Neighbors = append(copy.Neighbors, deepClone(neighbor, nodes))
	}

	return copy
}
