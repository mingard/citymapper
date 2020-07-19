// CityMapper Coding Test
// Arthur Mingard 2020

package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/mingard/citymapper/internal/app"
)

func main() {
	args := os.Args[1:]
	fmt.Println("ARGS", args)

	if len(args) != 3 {
		err := errors.New("Invalid arguements")
		log.Fatal(err)
		return
	}

	cm := app.New(args[0], args[1], args[2])
	cm.Run()
}
