package graph

import (
	"reflect"
	"testing"
)

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
	if g.verticesCount != 1 {
		t.Errorf("the count of edges expects 1, but have %v", g.verticesCount)
	}

	// sencond time to add a vertex 0, expects returnning error
	if err := g.AddVertex(0); err == nil {
		t.Error()
	}
}

func TestGetVertices(t *testing.T) {
	g := NewGraph()
	g.AddVertex(0)

	expect := []Vertex{0}
	actual := g.GetVertices()

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("GetVertices want: %v\nget: %v", expect, actual)
	}
}

func TestRemoveVertex(t *testing.T) {
	g := NewGraph()
	g.AddVertex(0)

	if err := g.RemoveVertex(0); err != nil {
		t.Error()
	}
	if g.verticesCount != 0 {
		t.Errorf("the count of edges expects 0, but have %v", g.verticesCount)
	}

	// expects error "doesn't exist"
	if err := g.RemoveVertex(0); err == nil {
		t.Error()
	}
}
