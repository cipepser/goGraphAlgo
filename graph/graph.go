package graph

// Vertex represents a vertex of the graph
type Vertex uint

// Edge represents a edge connects a vertex and another of the graph
type Edge struct {
	From Vertex
	To   Vertex
}

type graph struct {
	vertices   []Vertex
	edges      []Edge
	isDirected bool
}

// NewGraph constructs a new graph
func NewGraph() *graph {
	return new(graph)
}

func (g *graph) AddNode() error {

	return nil
}

func (g *graph) RemoveNode() error {

	return nil
}

func (g *graph) AddEdge() error {

	return nil
}

func (g *graph) RemoveEdge() error {

	return nil
}
