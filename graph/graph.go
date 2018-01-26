package graph

type Vertex uint

type Edge struct {
	From Vertex
	To   Vertex
}

type graph struct {
	isDirected bool
}

func NewGraph() *graph {
	return new(graph)
}
