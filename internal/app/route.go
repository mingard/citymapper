// CityMapper Coding Test
// Arthur Mingard 2020

package app

import (
	"context"
	"fmt"

	"github.com/mingard/citymapper/internal/pkg/dag"
)

type Route struct {
	datastore *dag.DAG
	distance  uint32
	start     *dag.Entry
	end       *dag.Entry
	current   *dag.Entry
	edges     map[uint32]bool
	ctx       context.Context
	cancel    context.CancelFunc
}

func (r *Route) Init() {
	for k, v := range r.current.Links {
		// fmt.Println(k, r.end.ID)
		if k == r.end.ID {
			fmt.Println("GOT TO THE END!!!", r.distance+v.Distance)
			fmt.Println("after", r.edges)
			continue
		}
		// fmt.Println(r.edges[k], k, r.start.ID)
		// This is a cyclical route or it links back to the origin. Skip it.
		if r.edges[k] || k == r.start.ID {
			// fmt.Println("CYCLICAL")
			return
		}

		next, err := r.datastore.Get(k)

		// No link found or link is cyclical.
		if err != nil {
			r.cancel()
			return
		}

		// Add the link distance to the routes current distance.
		dist := r.distance + v.Distance

		// Create an edge map including the next routes key.
		edges := make(map[uint32]bool, 0)

		// fmt.Println("before", edges)
		// Add all previous edges.
		for k, v := range r.edges {
			// fmt.Println("IS CALLED")
			edges[k] = v
		}

		edges[k] = true

		// fmt.Println(edges)
		// fmt.Println("after", edges)
		// fmt.Println(k, len(edges), dist, next)
		route := NewRoute(r.ctx, r.datastore, dist, r.start, next, r.end, edges)
		// fmt.Println("CARRY ON ", dist)

		// fmt.Println("Edges", len(route.edges))
		route.Init()

	}
}

func NewRoute(ctx context.Context, datastore *dag.DAG, distance uint32, start, current, end *dag.Entry, edges map[uint32]bool) *Route {
	ctx, cancel := context.WithCancel(ctx)
	return &Route{
		datastore: datastore,
		distance:  distance,
		start:     start,
		current:   current,
		end:       end,
		edges:     edges,
		ctx:       ctx,
		cancel:    cancel,
	}

}

// FindShortestRoute uses the massaged dataset to find the shortest route.
func (c *CityMapper) FindShortestRoute() error {
	start, err := c.datastore.Get(c.from)

	if err != nil {
		return err
	}

	end, err := c.datastore.Get(c.to)

	if err != nil {
		return err
	}

	route := NewRoute(c.ctx, c.datastore, uint32(0), start, start, end, make(map[uint32]bool, 0))

	route.Init()

	return nil
}
