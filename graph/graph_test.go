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
