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

func TestContains(t *testing.T) {
	s := NewVerSet()
	if s.Contains(0) == true {
		t.Error("s expects not to have any element, but get true")
	}

	s.Add(0)
	if s.Contains(0) == false {
		t.Error("s expects to have 0, but get false")
	}
}

func TestAdd(t *testing.T) {
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

func TestRemove(t *testing.T) {
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

func TestCardinality(t *testing.T) {
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

func TestEqual(t *testing.T) {
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

func TestDifference(t *testing.T) {
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

func TestIntersect(t *testing.T) {
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

func TestUnion(t *testing.T) {
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
