package set

import (
	"reflect"
	"strconv"
	"testing"
)

func TestNewIntSet(t *testing.T) {
	expect := IntSet{}

	actual := NewIntSet()

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
	}
}

func TestContains(t *testing.T) {
	s := NewIntSet()
	if s.Contains(0) == true {
		t.Error("s expects not to have any element, but get true")
	}

	s.Add(0)
	if s.Contains(0) == false {
		t.Error("s expects to have 0, but get false")
	}
}

func TestAdd(t *testing.T) {
	s := NewIntSet()
	if err := s.Add(0); err != nil {
		t.Error(err)
	}

	s.Add(0)
	if err := s.Add(0); err == nil {
		t.Error("expect to get error: 0 already exists")
	}

	expect := IntSet{
		0: struct{}{},
	}
	if !reflect.DeepEqual(expect, s) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, s)
	}
}

func TestRemove(t *testing.T) {
	s := NewIntSet()
	if err := s.Remove(0); err == nil {
		t.Error("expect to get error: 0 dose not exist")
	}

	s.Add(0)
	s.Add(1)
	if err := s.Remove(0); err != nil {
		t.Error(err)
	}

	expect := IntSet{
		1: struct{}{},
	}
	if !reflect.DeepEqual(expect, s) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, s)
	}
}

func TestCardinality(t *testing.T) {
	s := NewIntSet()

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

func TestDifference(t *testing.T) {
	// Difference(other Set) Set

}

func TestEqual(t *testing.T) {
	// Equal(other Set) bool

}

func TestIntersect(t *testing.T) {
	// Intersect(other Set) Set

}

func TestUnion(t *testing.T) {
	// Union(other Set) Set

}
