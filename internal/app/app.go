// CityMapper Coding Test
// Arthur Mingard 2020

package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mingard/citymapper/internal/pkg/dag"
)

// CityMapper is the application root instance.
type CityMapper struct {
	datastore  *dag.DAG
	sourceUrl  string
	from       uint32
	to         uint32
	distance   uint32
	totalNodes int64
	totalEdges int64
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
	c.sourceUrl = "https://s3-eu-west-1.amazonaws.com/citymapper-assets/citymapper-coding-test-graph.dat"
	// c.from = 876500321
	// c.to = 1524235806
	c.from = 1
	c.to = 6
}

// Run performs a lookup on the initialized data.
func (c *CityMapper) Run() {
	c.Must(c.FindShortestRoute())
}

// Initialize initializes all dependency classes.
func (c *CityMapper) Initialize() {
	defer c.Recover()

	c.InitializeStore()
	c.FetchData()
	// 	f.InitializeHTTP()
}

// HandleExit calls all exit methods before exiting the application.
func (c *CityMapper) HandleExit() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		c.cancel()
		// f.Must(f.State.Transition(StateShutdown))
	}()
}

// New returns a new FoundlandFrontend instance.
func New() *CityMapper {
	ctx, cancel := context.WithCancel(context.Background())
	cm := &CityMapper{
		ctx:    ctx,
		cancel: cancel,
	}

	cm.SetDefaults()
	fmt.Println("Start")
	cm.HandleExit()
	cm.Initialize()

	return cm
}
