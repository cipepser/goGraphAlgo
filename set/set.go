package set

import (
	"errors"
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

	// Difference returns the difference of the Set and {other}.
	Difference(other Set) Set

	// Equal checks whether {other} is same as the Set or not.
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

func (s IntSet) Remove(i int) error {

	return nil
}

func (s IntSet) Cardinality() int {

	return -1
}

func (s IntSet) Difference(other IntSet) IntSet {

	return nil
}

func (s IntSet) Equal(other IntSet) bool {

	return false
}

func (s IntSet) Intersect(other IntSet) IntSet {

	return nil
}

func (s IntSet) Union(other IntSet) IntSet {

	return nil
}
