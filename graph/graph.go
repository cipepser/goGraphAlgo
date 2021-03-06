package graph

import "errors"

// Vertex represents a vertex of the graph
type Vertex uint

// Edge represents a edge connects a vertex with another of the graph
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
	neighbours    map[Vertex][]Vertex
}

// NewGraph constructs a new graph
func NewGraph() *graph {
	g := new(graph)
	g.vertices = map[Vertex]int{}
	g.edges = map[Edge]int{}
	g.neighbours = map[Vertex][]Vertex{}
	return g
}

// SetDir sets `isDirected` to true or false by `dir`
func (g *graph) SetDir(dir bool) {
	g.isDirected = dir
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

// ExistsVertex chech whether the vertex `v` exists in the graph `g`, or not
func (g *graph) ExistsVertex(v Vertex) bool {
	if _, ok := g.vertices[v]; !ok {
		return false
	}

	return true
}

// AddVertex adds a vertex `v` to the graph `g`.
// If a graph has already have the vertex `v`, it returns an error
func (g *graph) AddVertex(v Vertex) error {
	if g.ExistsVertex(v) {
		return errors.New("g has already have the vertex v")
	}
	g.vertices[v]++
	g.verticesCount++
	return nil
}

// RemoveVertex remove a vertex form the graph `g`.
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

// ExistsEdge chech whether the edge difined by `from` and `to`
// exists in the graph `g`, or not
func (g *graph) ExistsEdge(from, to Vertex) bool {
	e := Edge{
		From: from,
		To:   to,
	}
	if _, ok := g.edges[e]; ok {
		return true
	}

	if !g.isDirected {
		e := Edge{
			From: to,
			To:   from,
		}
		if _, ok := g.edges[e]; ok {
			return true
		}
	}

	return false
}

// AddEdge adds an edge difined by `from` and `to` to the graph `g`.
// If a graph has already have the edge, it returns an error.
// The graph g must have the vertices `from` and `to`,
// if `g` don't have `from` or `to`, AddEdge returns an error.
// The vertices `from` and `to` must be different.
func (g *graph) AddEdge(from, to Vertex, weight int) error {
	if from == to {
		return errors.New("can not add ann edge of same vertex")
	}

	if !g.ExistsVertex(from) || !g.ExistsVertex(to) {
		return errors.New("edge doesn't exist")
	}

	if g.ExistsEdge(from, to) {
		return errors.New("edge already exists")
	}
	g.edges[Edge{
		From: from,
		To:   to,
	}] = weight

	g.neighbours[from] = append(g.neighbours[from], to)

	g.edgesCount++

	if !g.isDirected {
		g.neighbours[to] = append(g.neighbours[to], from)
	}

	return nil
}

// RemoveEdge removes an edge difined by `from` and `to`,
// form the graph `g`.
// The vertices `from` and `to` must be different.
func (g *graph) RemoveEdge(from, to Vertex) error {
	if from == to {
		return errors.New("can not remove an edge of same vertex")
	}

	if !g.ExistsEdge(from, to) {
		return errors.New("the edge doesn't exist")
	}

	delete(g.edges, Edge{
		From: from,
		To:   to,
	})

	g.edgesCount--
	return nil
}

// GetWeight returns a weight of the edge difined by `from` and `to`.
// The vertices `from` and `to` must be different.
// The edge must exist.
func (g *graph) GetWeight(from, to Vertex) (int, error) {
	if from == to {
		return 0, errors.New("can not get an edge of same vertex")
	}

	if !g.ExistsEdge(from, to) {
		return 0, errors.New("the edge doesn't exist")
	}

	weight := g.edges[Edge{
		From: from,
		To:   to,
	}]

	return weight, nil
}

// SetWeight sets a weight to the edge difined by `from` and `to`.
// The vertices `from` and `to` must be different.
// The edge must exist.
func (g *graph) SetWeight(from, to Vertex, weight int) error {
	if from == to {
		return errors.New("can not change an edge of same vertex")
	}

	if !g.ExistsEdge(from, to) {
		return errors.New("the edge doesn't exist")
	}

	g.edges[Edge{
		From: from,
		To:   to,
	}] = weight

	return nil
}

// GetNeighbours returns the neighbours of the Vertex `v`.
func (g *graph) GetNeighbours(v Vertex) []Vertex {
	vertices := make([]Vertex, g.vertices[v])

	return vertices
}
