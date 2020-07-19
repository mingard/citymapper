// CityMapper Coding Test
// Arthur Mingard 2020

package app

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const expectedEdgeParts = 3

// Close safely closes a stream.
func Close(c io.Closer) error {
	err := c.Close()
	if err != nil {
		return err
	}

	return nil
}

// doRequest performs an HTTP request
func doRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)

	return resp, err
}

// parseData converts the data to edges.
func (c *CityMapper) parseData(resp *http.Response) error {
	defer Close(resp.Body)
	scanner := bufio.NewScanner(resp.Body)

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

		// Break up the string
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

func (c *CityMapper) loadAndParseSource(url string) error {
	srcData, err := doRequest(url)

	// Error fetching data from source
	if err != nil {
		return err
	}

	c.parseData(srcData)

	return nil
}
