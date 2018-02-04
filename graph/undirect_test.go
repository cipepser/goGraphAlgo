package graph

import (
	"reflect"
	"strconv"
	"testing"
)

func TestNewVerSet(t *testing.T) {
	expect := VerSet{}

	actual := NewVerSet()

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
	}
}

func TestVerSetContains(t *testing.T) {
	s := NewVerSet()
	if s.Contains(0) == true {
		t.Error("s expects not to have any element, but get true")
	}

	s.Add(0)
	if s.Contains(0) == false {
		t.Error("s expects to have 0, but get false")
	}
}

func TestVerSetAdd(t *testing.T) {
	s := NewVerSet()
	if err := s.Add(0); err != nil {
		t.Error(err)
	}

	s.Add(0)
	if err := s.Add(0); err == nil {
		t.Error("expect to get error: 0 already exists")
	}

	expect := VerSet{
		0: struct{}{},
	}
	if !reflect.DeepEqual(expect, s) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, s)
	}
}

func TestVerSetRemove(t *testing.T) {
	s := NewVerSet()
	if err := s.Remove(0); err == nil {
		t.Error("expect to get error: 0 dose not exist")
	}

	s.Add(0)
	s.Add(1)
	if err := s.Remove(0); err != nil {
		t.Error(err)
	}

	expect := VerSet{
		1: struct{}{},
	}
	if !reflect.DeepEqual(expect, s) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, s)
	}
}

func TestVerSetCardinality(t *testing.T) {
	s := NewVerSet()

	s.Add(0)
	s.Add(1)

	if s.Cardinality() != 2 {
		t.Error("\nexpect: 2\nactual:", strconv.Itoa(s.Cardinality())+"\n")
	}

	s.Remove(0)
	if s.Cardinality() != 1 {
		t.Error("\nexpect: 1\nactual:", strconv.Itoa(s.Cardinality())+"\n")
	}
}

func TestVerSetEqual(t *testing.T) {
	s := NewVerSet()
	s.Add(0)
	s.Add(1)

	expect := VerSet{
		0: struct{}{},
		1: struct{}{},
	}
	if !s.Equal(expect) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, s)
	}
}

func TestVerSetDifference(t *testing.T) {
	s0 := NewVerSet()
	s0.Add(0)
	s0.Add(1)

	s1 := NewVerSet()
	s1.Add(1)

	expect := VerSet{
		0: struct{}{},
	}
	actual := s0.Difference(s1)
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
	}
}

func TestVerSetIntersect(t *testing.T) {
	s0 := NewVerSet()
	s0.Add(0)
	s0.Add(1)

	s1 := NewVerSet()
	s1.Add(1)

	expect := VerSet{
		1: struct{}{},
	}
	actual := s0.Intersect(s1)
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
	}
}

func TestVerSetUnion(t *testing.T) {
	s0 := NewVerSet()
	s0.Add(0)
	s0.Add(1)

	s1 := NewVerSet()
	s1.Add(1)
	s1.Add(2)

	expect := VerSet{
		0: struct{}{},
		1: struct{}{},
		2: struct{}{},
	}
	actual := s0.Union(s1)
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
	}
}

func TestDisjointSetAlgorithm(t *testing.T) {
	g := NewGraph()

	vertices := []Vertex{0, 1, 2, 3, 4, 5}
	for _, v := range vertices {
		g.AddVertex(v)
	}

	g.AddEdge(0, 1, 0)
	g.AddEdge(0, 3, 0)
	g.AddEdge(1, 2, 0)
	g.AddEdge(1, 4, 0)
	g.AddEdge(2, 5, 0)

	if DisjointSetAlgorithm(g) {
		t.Error("expect to get false, but get true")
	}

	g.AddEdge(3, 4, 0)

	if !DisjointSetAlgorithm(g) {
		t.Error("expect to get true, but get false")
	}
}

func TestNewDisjointSet(t *testing.T) {
	g := NewGraph()

	vertices := []Vertex{0, 1, 2, 3, 4, 5}
	for _, v := range vertices {
		g.AddVertex(v)
	}

	g.AddEdge(0, 1, 0)
	g.AddEdge(0, 3, 0)
	g.AddEdge(1, 2, 0)
	g.AddEdge(1, 4, 0)
	g.AddEdge(2, 5, 0)
	g.AddEdge(3, 4, 0)

	actual, err := g.NewDisjointSet()
	if err != nil {
		t.Error(err)
	}

	expect := DisjointSet{
		VerSet{0: struct{}{}},
		VerSet{1: struct{}{}},
		VerSet{2: struct{}{}},
		VerSet{3: struct{}{}},
		VerSet{4: struct{}{}},
		VerSet{5: struct{}{}},
	}

	if !actual.Equal(expect) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
	}
}

func TestDisjointSetEqual(t *testing.T) {
	expect := DisjointSet{
		VerSet{1: struct{}{}},
		VerSet{
			2: struct{}{},
			0: struct{}{},
		},
	}
	actual := DisjointSet{
		VerSet{
			0: struct{}{},
			2: struct{}{},
		},
		VerSet{1: struct{}{}},
	}

	if !actual.Equal(expect) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
	}

	actual = DisjointSet{
		VerSet{1: struct{}{}},
	}

	if actual.Equal(expect) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
	}
}

func TestDisjointSetFindSet(t *testing.T) {
	expectF := VerSet{
		2: struct{}{},
		0: struct{}{},
	}
	expectT := VerSet{1: struct{}{}}
	d := DisjointSet{expectF, expectT}

	e := Edge{
		From: 0,
		To:   1,
	}

	actualF, actualT := d.FindSet(e)

	if !reflect.DeepEqual(expectF, actualF) {
		t.Errorf("\nexpectF: %v\nactualF: %v\n", expectF, actualF)
	}
	if !reflect.DeepEqual(expectT, actualT) {
		t.Errorf("\nexpectT: %v\nactualT: %v\n", expectT, actualT)
	}
}

func TestDisjointSetUnion(t *testing.T) {
	expect := DisjointSet{
		VerSet{
			0: struct{}{},
			1: struct{}{},
			2: struct{}{},
		},
		VerSet{3: struct{}{}},
	}
	actual := DisjointSet{
		VerSet{
			2: struct{}{},
			0: struct{}{},
		},
		VerSet{1: struct{}{}},
		VerSet{3: struct{}{}},
	}
	actual = actual.Union(VerSet{1: struct{}{}}, VerSet{
		2: struct{}{},
		0: struct{}{},
	})

	if !actual.Equal(expect) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
	}
}
