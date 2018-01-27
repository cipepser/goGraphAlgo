package graph

import "testing"

func TestVisualize(t *testing.T) {
	g := NewGraph()
	g.SetDir(true)

	g.AddVertex(0)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)

	g.AddEdge(0, 1, 0)
	g.AddEdge(0, 2, 0)
	g.AddEdge(0, 3, 0)
	g.AddEdge(2, 3, 0)

	if err := g.Visualize(); err != nil {
		t.Error(err)
	}

}
