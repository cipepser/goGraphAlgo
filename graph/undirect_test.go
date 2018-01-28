package graph

import (
	"reflect"
	"testing"
)

func TestMakeSet(t *testing.T) {
	g := NewGraph()
	g.AddVertex(0)
	g.AddVertex(1)

	actual := MakeSet(g)

	expect := []map[Vertex]struct{}{
		map[Vertex]struct{}{0: struct{}{}},
		map[Vertex]struct{}{1: struct{}{}},
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expect: %v\nactual: %v", expect, actual)
	}
}
