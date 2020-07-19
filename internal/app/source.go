// CityMapper Coding Test
// Arthur Mingard 2020

package app

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/mingard/citymapper/internal/pkg/dag"
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

func doRequest(url string) (*http.Response, error) {
	resp, err := http.Get(url)

	return resp, err
}

func newLink(parts []string) (uint32, uint32, *dag.Link, error) {
	from, err := strconv.ParseUint(parts[0], 10, 32)

	if err != nil {
		return 0, 0, nil, err
	}
	to, err := strconv.ParseUint(parts[1], 10, 32)

	if err != nil {
		return 0, 0, nil, err
	}

	dist, err := strconv.ParseUint(parts[2], 10, 32)

	if err != nil {
		return 0, 0, nil, err
	}

	return uint32(from), uint32(to), &dag.Link{uint32(dist)}, nil
}

func (c *CityMapper) parseData(resp *http.Response) error {
	defer Close(resp.Body)
	scanner := bufio.NewScanner(resp.Body)
	fmt.Println("parseData")

	// Exit if line scanner fails.
	if err := scanner.Err(); err != nil {
		return err
	}

	// a := &Node{Name: "a"}
	// b := &Node{Name: "b"}
	// c := &Node{Name: "c"}
	// d := &Node{Name: "d"}
	// e := &Node{Name: "e"}
	// f := &Node{Name: "f"}
	// g := &Node{Name: "g"}

	nodeA := dag.NewEntry(uint32(1))
	nodeB := dag.NewEntry(uint32(2))
	nodeC := dag.NewEntry(uint32(3))
	nodeD := dag.NewEntry(uint32(4))
	nodeE := dag.NewEntry(uint32(5))
	nodeF := dag.NewEntry(uint32(6))
	nodeG := dag.NewEntry(uint32(7))

	nodeA.Links[3] = &dag.Link{uint32(2)}
	nodeA.Links[2] = &dag.Link{uint32(5)}
	nodeC.Links[2] = &dag.Link{uint32(1)}
	nodeC.Links[4] = &dag.Link{uint32(9)}
	nodeB.Links[4] = &dag.Link{uint32(4)}
	nodeD.Links[5] = &dag.Link{uint32(2)}
	nodeD.Links[7] = &dag.Link{uint32(30)}
	nodeD.Links[6] = &dag.Link{uint32(10)}
	nodeF.Links[7] = &dag.Link{uint32(1)}

	nodeC.Links[1] = &dag.Link{uint32(2)}
	nodeB.Links[1] = &dag.Link{uint32(5)}
	nodeB.Links[3] = &dag.Link{uint32(1)}
	nodeD.Links[3] = &dag.Link{uint32(9)}
	nodeD.Links[2] = &dag.Link{uint32(4)}
	nodeE.Links[4] = &dag.Link{uint32(2)}
	nodeG.Links[4] = &dag.Link{uint32(30)}
	nodeF.Links[4] = &dag.Link{uint32(10)}
	nodeG.Links[6] = &dag.Link{uint32(1)}

	c.datastore.Put(nodeA)
	c.datastore.Put(nodeB)
	c.datastore.Put(nodeC)
	c.datastore.Put(nodeD)
	c.datastore.Put(nodeE)
	c.datastore.Put(nodeF)
	c.datastore.Put(nodeG)

	// after map[2:true 3:true 4:true 6:true]
	// after map[b:true c:true d:true f:true]

	// graph := Graph{}
	// graph.AddEdge(a, c, 2)

	// graph.AddEdge(a, b, 5) First 5

	// graph.AddEdge(c, b, 1) Second 6

	// graph.AddEdge(c, d, 9) Third 15

	// graph.AddEdge(b, d, 4)

	// graph.AddEdge(d, e, 2) Fourth 17

	// graph.AddEdge(d, g, 30)

	// graph.AddEdge(d, f, 10) Fifth 27

	// graph.AddEdge(f, g, 1) End 28

	// Cycle through lines.
	// for scanner.Scan() {

	// 	// Set totalNodes when not set.
	// 	if c.totalNodes < 1 {
	// 		totalNodes, err := strconv.ParseInt(scanner.Text(), 10, 64)

	// 		// Capture parse failure
	// 		if err != nil {
	// 			return err
	// 		}

	// 		c.totalNodes = totalNodes
	// 		continue
	// 	}

	// 	// Create entries with the OSM ids.
	// 	if c.datastore.Size() < c.totalNodes {
	// 		id, err := strconv.ParseUint(scanner.Text(), 10, 32)

	// 		if err != nil {
	// 			return err
	// 		}

	// 		e := dag.NewEntry(uint32(id))

	// 		c.datastore.Put(e)
	// 		continue
	// 	}

	// 	if c.totalEdges < 1 {
	// 		totalEdges, err := strconv.ParseInt(scanner.Text(), 10, 64)

	// 		// Capture parse failure
	// 		if err != nil {
	// 			return err
	// 		}

	// 		c.totalEdges = totalEdges
	// 		continue
	// 	}

	// 	edgeParts := strings.Fields(scanner.Text())

	// 	if actual := len(edgeParts); actual != expectedEdgeParts {
	// 		return fmt.Errorf(`Failed to parse edge. Expected %d values. Got %d`, expectedEdgeParts, actual)
	// 	}

	// 	from, to, link, err := newLink(edgeParts)

	// 	if err != nil {
	// 		return err
	// 	}

	// 	// Create a reference to the related edge.
	// 	if edgeParent, err := c.datastore.Get(from); err == nil {
	// 		edgeParent.Links[to] = link
	// 	}

	// 	// Create a reverse reference from the related edge to the source.
	// 	if edgeParentReverse, err := c.datastore.Get(to); err == nil {
	// 		edgeParentReverse.Links[from] = link
	// 	}
	// }

	return nil
}

func (c *CityMapper) loadAndParseSource(url string) error {
	fmt.Println("loadAndParseSource")
	srcData, err := doRequest(url)

	// Error fetching data from source
	if err != nil {
		return err
	}

	c.parseData(srcData)

	return nil
}
