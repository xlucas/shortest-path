package path

// inifity represent an infinite distance
const infinity = -1

// Vertex represents a vertex within a graph.
type Vertex struct {
	Arcs     Arcs
	Distance int
	ID       string
	Prev     *Vertex
}

// NewVertex creates a new vertex with the specified ID and arcs to its neighbor
// vertexes.
func NewVertex(id string, a Arcs) *Vertex {
	return &Vertex{
		Arcs:     a,
		Distance: infinity,
		ID:       id,
	}
}

type vertexComparator struct{}

func (c *vertexComparator) Equal(i, j interface{}) bool {
	return i.(*Vertex).Distance == j.(*Vertex).Distance
}

func (c *vertexComparator) Less(i, j interface{}) bool {
	vi := i.(*Vertex)
	vj := j.(*Vertex)

	if vi.Distance == infinity && vj.Distance != infinity {
		return true
	}
	if vi.Distance != infinity && vj.Distance == infinity {
		return false
	}
	return vi.Distance > vj.Distance
}
