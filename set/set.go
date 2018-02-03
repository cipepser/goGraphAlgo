package set

import (
	"errors"
	"reflect"
	"strconv"
)

// Set is an interface should be implemented as a mathematically set.
type Set interface {
	// Add adds an `i` to the Set.
	Add(i interface{}) error

	// Remove removes an `i` from the Set.
	Remove(i interface{}) error

	// Cardinality returns the number of elements in the Set.
	Cardinality() int

	// Contains checks whether an `i` exists in the Set or not.
	Contains(i interface{}) bool

	// Difference returns the difference of the Set and `other`.
	Difference(other Set) Set

	// Equal checks whether `other` is same as the Set or not.
	Equal(other Set) bool

	// Intersect returns the Intersection of the Set and `other`.
	Intersect(other Set) Set

	// Union returns the Union of the Set and `other`
	Union(other Set) Set
}

// IntSet represents a set of ints.
type IntSet map[int]struct{}

// NewIntSet constructs a new IntSet.
func NewIntSet() IntSet {
	return make(IntSet, 0)
}

// Contains checks whether an `i` exists in `s` or not.
func (s IntSet) Contains(i int) bool {
	_, ok := s[i]
	return ok
}

// Add adds an `i` to `s`.
func (s IntSet) Add(i int) error {
	if s.Contains(i) {
		return errors.New(strconv.Itoa(i) + " already exists")
	}
	s[i] = struct{}{}
	return nil
}

// Remove removes an `i` from `s`.
func (s IntSet) Remove(i int) error {
	if !s.Contains(i) {
		return errors.New(strconv.Itoa(i) + " dose not exist")
	}
	delete(s, i)
	return nil
}

// Cardinality returns the number of elements in the Set.
func (s IntSet) Cardinality() int {
	return len(s)
}

// Equal checks whether `other` is same as `s` or not.
func (s IntSet) Equal(other IntSet) bool {
	return reflect.DeepEqual(s, other)
}

// Difference returns the difference of `s` and `other`.
// Difference have a referential transparency.
func (s IntSet) Difference(other IntSet) IntSet {
	d := NewIntSet()
	for i := range s {
		if _, ok := other[i]; !ok {
			d.Add(i)
		}
	}
	return d
}

// Intersect returns the Intersection of `s` and `other`.
// Intersect have a referential transparency.
func (s IntSet) Intersect(other IntSet) IntSet {
	inter := NewIntSet()
	for i := range s {
		if _, ok := other[i]; ok {
			inter.Add(i)
		}
	}
	return inter
}

// Union returns the Union of `s` and `other`
// Union have a referential transparency.
func (s IntSet) Union(other IntSet) IntSet {
	u := NewIntSet()
	for i := range s {
		u[i] = struct{}{}
	}
	for i := range other {
		u[i] = struct{}{}
	}

	return u
}
