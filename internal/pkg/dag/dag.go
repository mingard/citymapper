// CityMapper Coding Test
// Arthur Mingard 2020

package dag

import (
	"errors"
	"fmt"
	"sync"
)

type DAG struct {
	sync.RWMutex
	values sync.Map
	size   int64
}

// Size returns the total byte size of the DAG.
func (d *DAG) Size() int64 {
	return d.size
}

// find finds a value in the sync map and returns the entry.
func (d *DAG) find(k uint32) *Entry {
	if v, ok := d.values.Load(k); ok {
		e := v.(*Entry)
		return e
	}
	return nil
}

// Get retrieves a record by key.
func (d *DAG) Get(k uint32) (*Entry, error) {
	d.Lock()
	defer d.Unlock()

	if e := d.find(k); e != nil {
		return e, nil
	}

	return nil, fmt.Errorf("Unable to locate key %s", k)
}

// Put adds an entry.
func (d *DAG) Put(e *Entry) error {
	d.Lock()
	defer d.Unlock()

	if e.ID < 1 {
		return errors.New("Failed to insert Entry. ID is invalid")
	}

	// This is a new entry.
	if prev := d.find(e.ID); prev == nil {
		d.size++
	}

	d.values.Store(e.ID, e)

	return nil
}

// RangeAll calls an input method for each entry.
func (d *DAG) RangeAll(r func(k, v interface{}) bool) {
	d.values.Range(func(k, v interface{}) bool {
		return r(k, v)
	})
}

// New creates a new instance of DAG.
func New() *DAG {
	return new(DAG)
}
