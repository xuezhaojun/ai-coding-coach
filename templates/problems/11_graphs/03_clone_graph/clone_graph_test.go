package clone_graph

import "testing"

// buildGraph constructs a graph from an adjacency list and returns the node with Val=1.
func buildGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}
	nodes := make([]*Node, len(adjList)+1)
	for i := 1; i <= len(adjList); i++ {
		nodes[i] = &Node{Val: i}
	}
	for i, neighbors := range adjList {
		for _, n := range neighbors {
			nodes[i+1].Neighbors = append(nodes[i+1].Neighbors, nodes[n])
		}
	}
	return nodes[1]
}

// collectNodes does BFS to collect all nodes in the graph.
func collectNodes(node *Node) map[int]*Node {
	if node == nil {
		return nil
	}
	visited := map[int]*Node{}
	queue := []*Node{node}
	visited[node.Val] = node
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, n := range cur.Neighbors {
			if _, ok := visited[n.Val]; !ok {
				visited[n.Val] = n
				queue = append(queue, n)
			}
		}
	}
	return visited
}

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
	}{
		{
			name:    "four node cycle",
			adjList: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			name:    "single node no neighbors",
			adjList: [][]int{{}},
		},
		{
			name:    "nil graph",
			adjList: [][]int{},
		},
		{
			name:    "two connected nodes",
			adjList: [][]int{{2}, {1}},
		},
		{
			name:    "three node chain",
			adjList: [][]int{{2}, {1, 3}, {2}},
		},
		{
			name:    "star graph",
			adjList: [][]int{{2, 3, 4}, {1}, {1}, {1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adjList)
			cloned := cloneGraph(original)

			if original == nil {
				if cloned != nil {
					t.Errorf("expected nil, got non-nil clone")
				}
				return
			}

			origNodes := collectNodes(original)
			clonedNodes := collectNodes(cloned)

			if len(origNodes) != len(clonedNodes) {
				t.Errorf("node count mismatch: original=%d, cloned=%d", len(origNodes), len(clonedNodes))
				return
			}

			for val, oNode := range origNodes {
				cNode, ok := clonedNodes[val]
				if !ok {
					t.Errorf("cloned graph missing node %d", val)
					continue
				}
				if oNode == cNode {
					t.Errorf("node %d is the same pointer in original and clone", val)
				}
				if len(oNode.Neighbors) != len(cNode.Neighbors) {
					t.Errorf("node %d neighbor count mismatch: original=%d, cloned=%d",
						val, len(oNode.Neighbors), len(cNode.Neighbors))
				}
			}
		})
	}
}
