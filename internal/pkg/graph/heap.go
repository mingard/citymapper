// CityMapper Coding Test
// Arthur Mingard 2020

package graph

import hp "container/heap"

type path struct {
	value int
	nodes []string
}

type minPath []path

// Len exposes the length of the path nodes.
func (h minPath) Len() int {
	return len(h)
}

// Less compares compares two values to
// confirm whether the first is lower than the second.
func (h minPath) Less(i, j int) bool {
	return h[i].value < h[j].value
}

// Swap performs a node swap.
func (h minPath) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds a path to the heap.
func (h *minPath) Push(x interface{}) {
	*h = append(*h, x.(path))
}

// Push removes a path from the heap.
func (h *minPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type heap struct {
	values *minPath
}

func newHeap() *heap {
	return &heap{values: &minPath{}}
}

func (h *heap) push(p path) {
	hp.Push(h.values, p)
}

func (h *heap) pop() path {
	i := hp.Pop(h.values)
	return i.(path)
}
