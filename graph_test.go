package path

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestPath(t *testing.T) {
	graph := NewGraph(
		[]*Vertex{
			NewVertex("A", Arcs{
				{Dst: "B", Distance: 100},
				{Dst: "C", Distance: 50},
				{Dst: "D", Distance: 75},
			}),
			NewVertex("B", Arcs{{Dst: "E", Distance: 25}}),
			NewVertex("C", Arcs{{Dst: "E", Distance: 80}}),
			NewVertex("D", Arcs{{Dst: "F", Distance: 50}}),
			NewVertex("E", Arcs{{Dst: "G", Distance: 40}}),
			NewVertex("F", Arcs{{Dst: "G", Distance: 60}}),
			NewVertex("G", Arcs{}),
		},
	)
	path := graph.ShortestPath("A", "G")

	var pathStr string
	for _, v := range path {
		pathStr += v.ID
	}

	assert.Equal(t, "ABEG", pathStr)
}
