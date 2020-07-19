// CityMapper Coding Test
// Arthur Mingard 2020

package main

import "github.com/mingard/citymapper/internal/app"

func main() {
	// Open a channel to keep application alive.
	// done := make(chan bool)

	cm := app.New()
	cm.Run()

	// <-done
}
