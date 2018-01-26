package graph

import "errors"

// Vertex represents a vertex of the graph
type Vertex uint

// Edge represents a edge connects a vertex and another of the graph
type Edge struct {
	From Vertex
	To   Vertex
}

type graph struct {
	vertices      map[Vertex]int
	verticesCount int
	edges         []Edge
	edgesCount    int
	isDirected    bool
}

// NewGraph constructs a new graph
func NewGraph() *graph {
	g := new(graph)
	g.vertices = map[Vertex]int{}
	return g
}

// GetVertices returns slice of Vertices
func (g *graph) GetVertices() []Vertex {
	vertices := make([]Vertex, g.verticesCount)

	i := 0
	for v := range g.vertices {
		vertices[i] = v
		i++
	}

	return vertices
}

// AddVertex adds a vertex v to the graph g.
// If a graph has already have the vertex v, it returns an error
func (g *graph) AddVertex(v Vertex) error {
	if _, ok := g.vertices[v]; ok {
		return errors.New("g has already have the vertex v")
	}
	g.vertices[v]++
	g.verticesCount++
	return nil
}

// RemoveVertex remove a vertex form the graph g.
func (g *graph) RemoveVertex() error {

	return nil
}

func (g *graph) AddEdge() error {

	return nil
}

func (g *graph) RemoveEdge() error {

	return nil
}
