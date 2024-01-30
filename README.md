# Citymapper Router Challenge

This application uses [Dijkstra's Algorithm](https://en.wikipedia.org/wiki/Dijkstra's_algorithm) to calculate the shortest distance between two OSM Nodes, using OSM graph data. It only uses go standard library.

## Requirements

This build requires [go 1.13](https://golang.org/dl/#go1.13.14) to be installed and configured according to the official [go installation instructions](https://golang.org/doc/install)

## Compiling the build

I have included a Makefile to handle build, which will compile the binary automatically before running the application.

If you wish to build it manually, run `make build`

## Running the application

As requested, this application must be run using the following format.

```
./run.sh <path-to-graph> <from-osm-id> <to-osm-id>
```

Example:

```shell
./run.sh citymapper-coding-test-graph.dat 876500321 1524235806
```

### Other Makefile options

`make clean` will remove any previously cached builds.

`make version` will output the version info.

`make fmt` will format the code to [gofmt](https://golang.org/cmd/gofmt/) standards.

## Design decisions

I used the Go's [heap](https://golang.org/pkg/container/heap/) as an efficient way to store, prioritise and filter edge paths.

Whilst the OSM graph included both Nodes and Edges, I did not explicitly store Nodes. Instead, the application creates Nodes during the Edge import stage, by checking Edge _from_ and _to_ for their existence in the Node map, and creating if missing. When benchmarking, this was a significant performance improvement against reading the scanned Node lines.

There were no asynchronous read/write operations against any of the data stores, or go routines, I avoided using locks, waitgroups or sync maps. If the data was ephemeral, transient or mutable I would've certainly used more rugged methods to mitigate pointer errors.

## Further Ideas

Theoretically it would be possible to start the process from both origin and destination and close a path at the intersection, which would've been my next benchmark.

I originally started using a Directed Acyclic Graph and stored the addresses as uint32. I removed this method during experimentation with Dijkstra's Algorithm, and instead used a heap, but it would be interesting the benchmark the performance as uint32 have far lighter memory _and_ CPU usage than string pointers.
