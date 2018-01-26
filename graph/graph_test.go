package graph

import "testing"

func TestIsDirected(t *testing.T) {
	expected := false

	g := NewGraph()
	actual := g.isDirected

	if expected != actual {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}

func TestAddVertex(t *testing.T) {
	g := NewGraph()

	// first time to add a vertex 0
	if err := g.AddVertex(0); err != nil {
		t.Error()
	}

	// sencond time to add a vertex 0, expects returnning error
	if err := g.AddVertex(0); err == nil {
		t.Error()
	}
}
