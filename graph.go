package path

import (
	"sort"

	"github.com/xlucas/heap"
)

// Graph represents an arranged set of vertexes.
type Graph struct {
	remaining *heap.Heap
	vertexMap map[string]*Vertex
}

// NewGraph creates a graph from a collection of vertexes.
func NewGraph(vertexes []*Vertex) *Graph {
	return &Graph{
		vertexMap: prepareMap(vertexes),
	}
}

// pathTo traces back the path computed by the shortest path algorithm and
// returns it as an order list of vertexes, starting from origin, towards
// the destination.
func (g *Graph) pathTo(dst string) []*Vertex {
	var path []*Vertex
	for vertex := g.vertexMap[dst]; vertex != nil; vertex = vertex.prev {
		path = append([]*Vertex{vertex}, path...)
	}
	return path
}

// prepareOrigin is used to set the origin's distance to 0 and its neighbors
// distances to the distance of each arc that leads to them.
func (g *Graph) prepareOrigin(origin string) {
	src := g.vertexMap[origin]
	src.distance = 0
	for _, arc := range src.Arcs {
		g.vertexMap[arc.Dst].distance = arc.Distance
		g.vertexMap[arc.Dst].prev = src
	}
}

// prepareRemainingSet is used to create the set of vertexes that needs to be
// visited by the shortest path algorithm.
func (g *Graph) prepareRemainingSet() {
	var slice []interface{}
	for _, v := range g.vertexMap {
		slice = append(slice, v)
	}
	g.remaining = heap.Heapify(slice, new(vertexComparator))
}

// ShortestPath computes the shortest path from src to dst within the graph.
func (g *Graph) ShortestPath(src, dst string) []*Vertex {
	g.prepareOrigin(src)
	g.prepareRemainingSet()

	// The first vertex is the origin of the graph.
	minVertex := g.remaining.Pop().(*Vertex)

	for minVertex != nil {
		minVertex = g.remaining.Pop().(*Vertex)

		if minVertex.ID == dst {
			return g.pathTo(dst)
		}

		sort.Sort(minVertex.Arcs)
		for _, arc := range minVertex.Arcs {
			dist := minVertex.distance + arc.Distance
			if neighbor := g.vertexMap[arc.Dst]; neighbor.distance == infinity || dist < neighbor.distance {
				neighbor.distance = dist
				neighbor.prev = minVertex
			}
		}

	}

	return nil
}

func prepareMap(vertexes []*Vertex) map[string]*Vertex {
	m := make(map[string]*Vertex)
	for _, v := range vertexes {
		m[v.ID] = v
	}
	return m
}
