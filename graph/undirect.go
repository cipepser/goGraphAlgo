package graph

import (
	"errors"
	"reflect"
	"strconv"
)

// VerSet represents a set of vertices.
type VerSet map[Vertex]struct{}

// NewVerSet constructs a new VerSet.
func NewVerSet() VerSet {
	return make(VerSet, 0)
}

// Contains checks whether an `v` exists in `s` or not.
func (s VerSet) Contains(v Vertex) bool {
	_, ok := s[v]
	return ok
}

// Add adds an `v` to `s`.
func (s VerSet) Add(v Vertex) error {
	if s.Contains(v) {
		return errors.New(strconv.FormatUint(uint64(v), 10) + " already exists")
	}
	s[v] = struct{}{}
	return nil
}

// Remove removes an `v` from `s`.
func (s VerSet) Remove(v Vertex) error {
	if !s.Contains(v) {
		return errors.New(strconv.FormatUint(uint64(v), 10) + " dose not exist")
	}
	delete(s, v)
	return nil
}

// Cardinality returns the number of elements in the Set.
func (s VerSet) Cardinality() int {
	return len(s)
}

// Equal checks whether `other` is same as `s` or not.
func (s VerSet) Equal(other VerSet) bool {
	return reflect.DeepEqual(s, other)
}

// Difference returns the difference of `s` and `other`.
// Difference have a referential transparency.
func (s VerSet) Difference(other VerSet) VerSet {
	d := NewVerSet()
	for v := range s {
		if _, ok := other[v]; !ok {
			d.Add(v)
		}
	}
	return d
}

// Intersect returns the Intersection of `s` and `other`.
// Intersect have a referential transparency.
func (s VerSet) Intersect(other VerSet) VerSet {
	inter := NewVerSet()
	for v := range s {
		if _, ok := other[v]; ok {
			inter.Add(v)
		}
	}
	return inter
}

// Union returns the Union of `s` and `other`
// Union have a referential transparency.
func (s VerSet) Union(other VerSet) VerSet {
	u := NewVerSet()
	for v := range s {
		u[v] = struct{}{}
	}
	for v := range other {
		u[v] = struct{}{}
	}

	return u
}

//////////////////////////////////////////////////////////////////

// Family is set of VerSet.
type Family map[*VerSet]struct{}

// NewFamily constructs a new Family.
func NewFamily() Family {
	return make(Family, 0)
}

// // Contains checks whether an `v` exists in f or not.
// func (f Family) Contains(s VerSet) bool {
// 	_, ok := f[s]
// 	return ok
// }

// // Add adds an `v` to f.
// func (f Family) Add(s VerSet) error {
// 	if f.Contains(s) {
// 		return errors.New(strconv.FormatUint(uint64(s), 10) + " already exists")
// 	}
// 	f[s] = struct{}{}
// 	return nil
// }
//
// // Remove removes an `v` from f.
// func (f Family) Remove(s VerSet) error {
// 	if !f.Contains(s) {
// 		return errors.New(strconv.FormatUint(uint64(s), 10) + " dose not exist")
// 	}
// 	delete(f, s)
// 	return nil
// }
//
// // Cardinality returns the number of elements in the Set.
// func (f Family) Cardinality() int {
// 	return len(f)
// }
//
// // Equal checks whether `other` is same as f or not.
// func (f Family) Equal(other Family) bool {
// 	return reflect.DeepEqual(f, other)
// }
//
// // Difference returns the difference of f and `other`.
// // Difference have a referential transparency.
// func (f Family) Difference(other Family) Family {
// 	d := NewFamily()
// 	for s := range f {
// 		if _, ok := other[s]; !ok {
// 			d.Add(s)
// 		}
// 	}
// 	return d
// }
//
// // Intersect returns the Intersection of f and `other`.
// // Intersect have a referential transparency.
// func (f Family) Intersect(other Family) Family {
// 	inter := NewFamily()
// 	for s := range f {
// 		if _, ok := other[s]; ok {
// 			inter.Add(s)
// 		}
// 	}
// 	return inter
// }
//
// // Union returns the Union of f and `other`
// // Union have a referential transparency.
// func (f Family) Union(other Family) Family {
// 	u := NewFamily()
// 	for s := range f {
// 		u[s] = struct{}{}
// 	}
// 	for s := range other {
// 		u[s] = struct{}{}
// 	}
//
// 	return u
// }
