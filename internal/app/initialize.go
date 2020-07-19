// CityMapper Coding Test
// Arthur Mingard 2020

package app

import (
	"github.com/mingard/citymapper/internal/pkg/dag"
)

// InitializeStore creates a new Directed acyclic graph store.
func (c *CityMapper) InitializeStore() {
	c.datastore = dag.New()
}

// FetchData retreives the OSM data from the source.
func (c *CityMapper) FetchData() {
	c.Must(c.loadAndParseSource(c.sourceUrl))
}
