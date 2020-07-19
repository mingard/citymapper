// CityMapper Coding Test
// Arthur Mingard 2020

package app

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const expectedEdgeParts = 3

// parseData converts the data to edges.
func (c *CityMapper) parseData(r io.Reader) error {
	scanner := bufio.NewScanner(r)

	// Exit if line scanner fails.
	if err := scanner.Err(); err != nil {
		return err
	}

	// Cycle through lines.

	counter, edgeStartPos := 0, 0
	for scanner.Scan() {
		// Inc counter first.
		counter++

		// First entry stores the total Nodes count.
		if counter == 1 {
			tn, err := strconv.Atoi(scanner.Text())

			// Capture parse failure
			if err != nil {
				return err
			}

			// Add offset to account for Node and Edge count rows.
			edgeStartPos = tn + 3
			continue
		}

		// continue loop until we reach the first edge.
		if counter < edgeStartPos {
			continue
		}

		// Break up the string.
		parts := strings.Fields(scanner.Text())

		if actual := len(parts); actual != expectedEdgeParts {
			return fmt.Errorf(`Failed to parse edge. Expected %d values. Got %d`, expectedEdgeParts, actual)
		}

		dist, err := strconv.Atoi(parts[2])

		if err != nil {
			return err
		}

		c.graph.AddEdge(parts[0], parts[1], dist)

		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CityMapper) loadAndParseSource() error {
	file, err := os.Open(c.sourcePath)
	if err != nil {
		return err
	}

	// Error fetching the data from source.
	if err != nil {
		return err
	}

	c.parseData(file)

	return nil
}
