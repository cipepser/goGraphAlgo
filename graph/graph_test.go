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

func TestExistsVertex(t *testing.T) {
	g := NewGraph()

	if g.ExistsVertex(0) {
		t.Error("vertex exists, expects existing")
	}

	g.AddVertex(0)
	if !g.ExistsVertex(0) {
		t.Error("vertex does not exists, expects not existing")
	}
}

func TestAddVertex(t *testing.T) {
	g := NewGraph()

	// first time to add a vertex 0
	if err := g.AddVertex(0); err != nil {
		t.Error(err)
	}
	if g.verticesCount != 1 {
		t.Errorf("the count of edges expects 1, but have %v", g.verticesCount)
	}

	// sencond time to add a vertex 0, expects returnning error
	if err := g.AddVertex(0); err == nil {
		t.Error(err)
	}
}

func TestGetVertices(t *testing.T) {
	g := NewGraph()

	if !reflect.DeepEqual(g.GetVertices(), []Vertex{}) {
		t.Error("expect empty slice")
	}

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
		t.Error(err)
	}
	if g.verticesCount != 0 {
		t.Errorf("the count of vertices expects 0, but have %v", g.verticesCount)
	}

	// expects error "doesn't exist"
	if err := g.RemoveVertex(0); err == nil {
		t.Error(err)
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph()

	if err := g.AddEdge(0, 1, 0); err == nil {
		t.Error("expect to get error: edge doesn't exist")
	}

	if g.edgesCount != 0 {
		t.Errorf("the count of edges expects 0, but have %v", g.edgesCount)
	}

	g.AddVertex(0)
	if err := g.AddEdge(0, 1, 0); err == nil {
		t.Error("expect to get error: edge doesn't exist")
	}

	if err := g.AddEdge(0, 0, 0); err == nil {
		t.Error("expect to get error: can not add edge to same vertex")
	}

	g.AddVertex(0)
	g.AddVertex(1)
	if err := g.AddEdge(0, 1, 0); err != nil {
		t.Error(err)
	}
	if g.edgesCount != 1 {
		t.Errorf("the count of edges expects 1, but have %v", g.edgesCount)
	}

	if err := g.AddEdge(0, 1, 0); err == nil {
		t.Error("expect to get error: edge already exists")
	}
}

func TestExistsEdge(t *testing.T) {
	g := NewGraph()
	if g.ExistsEdge(0, 1) {
		t.Error("expect to get false, the edge doesn't exist")
	}

	g.AddVertex(0)
	g.AddVertex(1)
	g.AddEdge(0, 1, 0)
	if !g.ExistsEdge(0, 1) {
		t.Error("expect to get true, the edge exists")
	}

}

func TestGetEdge(t *testing.T) {
	g := NewGraph()

	if !reflect.DeepEqual(g.GetEdges(), []Edge{}) {
		t.Error("expect empty slice")
	}

	g.AddVertex(0)
	g.AddVertex(1)
	g.AddEdge(0, 1, 0)

	expect := []Edge{Edge{
		From: 0,
		To:   1,
	}}
	actual := g.GetEdges()

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("GetVertices want: %v\nget: %v", expect, actual)
	}
}

func TestRemoveEdge(t *testing.T) {
	g := NewGraph()
	g.AddVertex(0)
	g.AddVertex(1)
	g.AddEdge(0, 1, 0)

	if err := g.RemoveEdge(0, 0); err == nil {
		t.Error("expect to get error: can not remove an edge of same vertex")
	}

	if err := g.RemoveEdge(0, 2); err == nil {
		t.Error("expect to get error: the edge doesn't exist")
	}

	if err := g.RemoveEdge(0, 1); err != nil {
		t.Error(err)
	}
	if g.edgesCount != 0 {
		t.Errorf("the count of vertices expects 0, but have %v", g.edgesCount)
	}
}
