package set

import (
	"reflect"
	"testing"
)

func TestNewIntSet(t *testing.T) {
	expect := IntSet{}

	actual := NewIntSet()

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("\nexpect: %v\nactual: %v\n", expect, actual)
	}
}

func TestAdd(t *testing.T) {
	// Add(i interface{}) error

}

func TestRemove(t *testing.T) {
	// Remove(i interface{}) error

}

func TestCardinality(t *testing.T) {
	// Cardinality() int

}

func TestContains(t *testing.T) {
	// Contains(i interface{}) bool

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
