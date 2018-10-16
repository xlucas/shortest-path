package path

// Arc represents an arc between two vertexes in a graph.
type Arc struct {
	Dst      string
	Distance int
}

// Arcs is a slice of arcs.
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
