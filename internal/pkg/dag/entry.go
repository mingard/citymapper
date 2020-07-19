// CityMapper Coding Test
// Arthur Mingard 2020

package dag

import (
	"sync"
)

// Link is a graph edge link definition.
type Link struct {
	Distance uint32
}

// Entry stores entry data.
type Entry struct {
	sync.RWMutex
	ID    uint32
	Links map[uint32]*Link
}

// NewEntry creates a new instance of Entry.
func NewEntry(id uint32) *Entry {
	return &Entry{
		ID:    id,
		Links: make(map[uint32]*Link, 0),
	}
}
