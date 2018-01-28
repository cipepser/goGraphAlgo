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

func TestUnion(t *testing.T) {
	A := map[Vertex]struct{}{0: struct{}{}}
	B := map[Vertex]struct{}{1: struct{}{}}

	actual := Union(A, B)

	expect := map[Vertex]struct{}{
		0: struct{}{},
		1: struct{}{},
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("expect: %v\nactual: %v", expect, actual)
	}
}

func TestFindSet(t *testing.T) {
	g := NewGraph()
	g.AddVertex(0)
	g.AddVertex(1)
	g.AddEdge(0, 1, 0)

	e := Edge{From: 0, To: 1}
	U := MakeSet(g)

	actA, actB := FindSet(e, U)

	expA := map[Vertex]struct{}{0: struct{}{}}
	expB := map[Vertex]struct{}{1: struct{}{}}

	if !reflect.DeepEqual(expA, actA) || !reflect.DeepEqual(expB, actB) {
		t.Errorf("\n[expect]\nA: %v\nB: %v\n[actual]\nA: %v\nB: %v",
			expA, expB, actA, actB)
	}

}
