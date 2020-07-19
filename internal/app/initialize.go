// CityMapper Coding Test
// Arthur Mingard 2020

package app

import (
	"github.com/mingard/citymapper/internal/pkg/graph"
)

// InitializeGraph creates a new graph instance.
func (c *CityMapper) InitializeGraph() {
	c.graph = graph.New()
}

// FetchData retreives the OSM data from the source.
func (c *CityMapper) FetchData() {
	c.Must(c.loadAndParseSource(c.sourceUrl))
}
