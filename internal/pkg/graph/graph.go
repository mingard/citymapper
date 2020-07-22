// CityMapper Coding Test
// Arthur Mingard 2020

package graph

// edge contains details of a related node and the distance (weight).
type edge map[string]int

// Graph stores graph nodes.
type Graph struct {
	nodes map[string]edge
}

// AddEdge creates a reference and reverse reference edge.
func (g *Graph) AddEdge(origin, dest string, weight int) {
	// Create nodes if missing.
	if g.nodes[origin] == nil {
		g.nodes[origin] = make(edge)
	}
	if g.nodes[dest] == nil {
		g.nodes[dest] = make(edge)
	}

	g.nodes[origin][dest] = weight
	g.nodes[dest][origin] = weight
}

// GetPath locates the shortest path between two nodes.
func (g *Graph) GetPath(origin, dest string) (int, []string) {
	// Create an empty heap, then push our origin node into it.
	h := newHeap()
	h.push(path{value: 0, nodes: []string{origin}})
	// Empty store of visited nodes.
	visited := make(map[string]bool)

	// This loop will continue as long as we keep pushing nodes into the heap.
	// Initially, it would only run once.
	for len(*h.values) > 0 {
		// Find the closest node we haven't attempted.
		// Pop returns and removes a node from the heap.
		p := h.pop()
		// Get the last edge in the node list.
		node := p.nodes[len(p.nodes)-1]

		// Skip if this node has already been attempted.
		if visited[node] {
			continue
		}

		// If the node is the destination exit the loop with a return.
		if node == dest {
			return p.value, p.nodes
		}
		// For each edge in the corresponding node in the graphs node list.
		for edge, weight := range g.nodes[node] {
			// Skip if we've visited this edge.
			if !visited[edge] {
				// Push our edge into the node list. 
				// When the loop cycles again this will be our new parent node.
				nodes := append([]string{}, append(p.nodes, edge)...)
				// Increase and store the total weight, and update the list of visited nodes.
				h.push(path{value: p.value + weight, nodes: nodes})
			}
		}

		// Add this node to the list of visited nodes.
		visited[node] = true
	}

	return 0, nil
}

// New creates a new instance of graph
func New() *Graph {
	return &Graph{
		nodes: make(map[string]edge),
	}
}
