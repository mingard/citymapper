// CityMapper Coding Test
// Arthur Mingard 2020

package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mingard/citymapper/internal/pkg/graph"
)

// CityMapper is the application root instance.
type CityMapper struct {
	graph      *graph.Graph
	sourcePath string
	from       string
	to         string
	ctx        context.Context
	cancel     context.CancelFunc
}

// Must handles errors.
func (c *CityMapper) Must(err error) {
	if err != nil {
		fmt.Printf("Must error: %v", err)
	}
}

// Recover is used to recover from panic attacks.
func (c *CityMapper) Recover() {
	if err := recover(); err != nil {
		fmt.Printf("Recovered from panic: %v", err)
	}
}

// SetDefaults sets the default configuration values.
func (c *CityMapper) SetDefaults() {
	c.sourcePath = "citymapper-coding-test-graph.dat"
	c.from = "876500321"
	c.to = "1524235806"
}

// Run performs a lookup on the initialized data.
func (c *CityMapper) Run() {
	dist, _ := c.graph.GetPath(c.from, c.to)
	fmt.Println(dist)
}

// Initialize initializes all dependency classes.
func (c *CityMapper) Initialize() {
	defer c.Recover()

	c.InitializeGraph()
	c.FetchData()
}

// HandleExit calls all exit methods before exiting the application.
func (c *CityMapper) HandleExit() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		c.cancel()
	}()
}

// New returns a new FoundlandFrontend instance.
func New(source, from, to string) *CityMapper {
	ctx, cancel := context.WithCancel(context.Background())
	cm := &CityMapper{
		ctx:    ctx,
		cancel: cancel,
	}

	cm.SetDefaults()

	cm.sourcePath = source
	cm.from = from
	cm.to = to

	cm.HandleExit()
	cm.Initialize()

	return cm
}
