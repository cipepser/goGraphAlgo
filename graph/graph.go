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
	edges         map[Edge]int
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

// ExistsVertex chech whether the vertex v exists in the graph g, or not
func (g *graph) ExistsVertex(v Vertex) bool {
	if _, ok := g.vertices[v]; !ok {
		return false
	}

	return true
}

// AddVertex adds a vertex v to the graph g.
// If a graph has already have the vertex v, it returns an error
func (g *graph) AddVertex(v Vertex) error {
	if g.ExistsVertex(v) {
		return errors.New("g has already have the vertex v")
	}
	g.vertices[v]++
	g.verticesCount++
	return nil
}

// RemoveVertex remove a vertex form the graph g.
func (g *graph) RemoveVertex(v Vertex) error {
	if !g.ExistsVertex(v) {
		return errors.New("input v doesn't exist in g")
	}

	delete(g.vertices, v)
	g.verticesCount--
	return nil
}

// GetEdges returns slice of Edges
func (g *graph) GetEdges() []Edge {
	edges := make([]Edge, g.edgesCount)

	i := 0
	for edge := range g.edges {
		edges[i] = edge
		i++
	}

	return edges
}

func (g *graph) ExistsEdge(from, to Vertex) bool {

	return false
}

func (g *graph) AddEdge(from, to Vertex, weight int) error {
	if from == to {
		return errors.New("can not add edge to same vertex")
	}

	if !g.ExistsVertex(from) || !g.ExistsVertex(to) {
		return errors.New("edge doesn't exist")
	}

	g.edges[Edge{
		From: from,
		To:   to,
	}] = weight

	g.edgesCount++

	return nil
}

func (g *graph) RemoveEdge() error {

	return nil
}
