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

	expect := []Set{
		Set{0: struct{}{}},
		Set{1: struct{}{}},
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("\nexpect: %v\nactual: %v", expect, actual)
	}
}

func TestUnion(t *testing.T) {
	A := Set{0: struct{}{}}
	B := Set{1: struct{}{}}

	U := []Set{
		Set{0: struct{}{}},
		Set{1: struct{}{}},
		Set{2: struct{}{}},
	}

	actC, actU := Union(A, B, U)

	expC := Set{
		0: struct{}{},
		1: struct{}{},
	}
	expU := []Set{
		expC,
		Set{2: struct{}{}},
	}

	if !reflect.DeepEqual(expC, actC) {
		t.Errorf("expected C and actual C is different\nexpect: %v\nactual: %v", expC, actC)
	}
	if !reflect.DeepEqual(expU, actU) {
		t.Errorf("expected U and actual U is different\nexpect: %v\nactual: %v", expU, actU)
	}
}

func TestContains(t *testing.T) {
	A := Set{0: struct{}{}}

	U := []Set{
		Set{0: struct{}{}},
		Set{1: struct{}{}},
		Set{2: struct{}{}},
	}

	if !Contains(A, U) {
		t.Error("U is epected to contains A, but not\n")
	}

	B := Set{3: struct{}{}}
	if Contains(B, U) {
		t.Error("U is epected not to contains B, but contains\n")
	}
}

func TestAdd(t *testing.T) {
	A := Set{0: struct{}{}}

	U := []Set{
		Set{0: struct{}{}},
		Set{1: struct{}{}},
	}
	if _, err := Add(A, U); err == nil {
		t.Error("expect to get error: A is aleady in U")
	}

	B := Set{2: struct{}{}}

	actual, err := Add(B, U)
	if err != nil {
		t.Error(err)
	}

	expect := []Set{
		Set{0: struct{}{}},
		Set{1: struct{}{}},
		Set{2: struct{}{}},
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
	}
}

func TestRemove(t *testing.T) {
	A := Set{3: struct{}{}}

	U := []Set{
		Set{0: struct{}{}},
		Set{1: struct{}{}},
		Set{2: struct{}{}},
	}
	if _, err := Remove(A, U); err == nil {
		t.Error("expect to get error: A does not exist in U")
	}

	B := Set{1: struct{}{}}

	actual, err := Remove(B, U)
	if err != nil {
		t.Error(err)
	}

	expect := []Set{
		Set{0: struct{}{}},
		Set{2: struct{}{}},
	}

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
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

	expA := Set{0: struct{}{}}
	expB := Set{1: struct{}{}}

	if !reflect.DeepEqual(expA, actA) || !reflect.DeepEqual(expB, actB) {
		t.Errorf("\n[expect]\nA: %v\nB: %v\n[actual]\nA: %v\nB: %v",
			expA, expB, actA, actB)
	}

}
