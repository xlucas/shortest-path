package path

// Arc represents a path between two vertexes in a graph.
type Arc struct {
	Dst      string
	Distance int
}

// Arcs is used to represent a slice of arc and makes it easier to sort them by
// implementing the sort interface.
type Arcs []*Arc

func (a Arcs) Len() int {
	return len(a)
}

func (a Arcs) Less(i, j int) bool {
	return a[i].Distance < a[j].Distance
}

func (a Arcs) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
