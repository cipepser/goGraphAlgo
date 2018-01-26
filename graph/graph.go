package graph

type Vertex uint

type Edge struct {
	From Vertex
	To   Vertex
}

type graph struct {
	isDirected bool
}

// NewGraph constructs a new graph
func NewGraph() *graph {
	return new(graph)
}
