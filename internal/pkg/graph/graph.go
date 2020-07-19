// CityMapper Coding Test
// Arthur Mingard 2020

package graph

// edge contains details of a related node and the distance (weight).
type edge struct {
	node   string
	weight int
}

// Graph stores graph nodes.
type Graph struct {
	nodes map[string][]edge
}

// AddEdge creates a reference and reverse reference edge.
func (g *Graph) AddEdge(origin, destiny string, weight int) {
	g.nodes[origin] = append(g.nodes[origin], edge{
		node:   destiny,
		weight: weight,
	})
	g.nodes[destiny] = append(g.nodes[destiny], edge{
		node:   origin,
		weight: weight,
	})
}

// GetEdges fetches a node from the data store.
func (g *Graph) GetEdges(node string) []edge {
	return g.nodes[node]
}

// GetPath locates the shortest path between two nodes.
func (g *Graph) GetPath(origin, destiny string) (int, []string) {
	h := newHeap()
	h.push(path{value: 0, nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		// Skip if this node has already been visited.
		if visited[node] {
			continue
		}

		if node == destiny {
			return p.value, p.nodes
		}

		for _, e := range g.GetEdges(node) {
			if !visited[e.node] {
				// Increase and store the total weight, and push the latest path node.
				h.push(path{value: p.value + e.weight, nodes: append([]string{}, append(p.nodes, e.node)...)})
			}
		}

		visited[node] = true
	}

	return 0, nil
}

// New creates a new instance of graph
func New() *Graph {
	return &Graph{
		nodes: make(map[string][]edge),
	}
}
