package path

// inifity represent an infinite distance
const infinity = -1

// Vertex represents a vertex within a graph.
type Vertex struct {
	Arcs     Arcs
	ID       string
	distance int
	prev     *Vertex
}

// NewVertex creates a new vertex with the specified ID and arcs to its neighbor
// vertexes.
func NewVertex(id string, a Arcs) *Vertex {
	return &Vertex{
		Arcs:     a,
		distance: infinity,
		ID:       id,
	}
}

type vertexComparator struct{}

func (c *vertexComparator) Equal(i, j interface{}) bool {
	return i.(*Vertex).distance == j.(*Vertex).distance
}

func (c *vertexComparator) Less(i, j interface{}) bool {
	vi := i.(*Vertex)
	vj := j.(*Vertex)

	if vi.distance == infinity && vj.distance != infinity {
		return true
	}
	if vi.distance != infinity && vj.distance == infinity {
		return false
	}
	return vi.distance > vj.distance
}
