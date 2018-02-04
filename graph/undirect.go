package graph

import (
	"errors"
	"reflect"
	"sort"
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

// ExistsCycle check whether a cycle exists or not.
// You can specify the algorithm by method:
// "DisjointSet": Disjoint-Set Algorithm
func (g *graph) ExistsCycle(method string) bool {
	if method == "DisjointSet" {
		return DisjointSetAlgorithm(g)
	}
	return false
}

// DisjointSetAlgorithm check whether a cycle exists or not.
func DisjointSetAlgorithm(g *graph) bool {
	d, err := g.NewDisjointSet()
	if err != nil {
		panic(err)
	}

	for _, e := range g.GetEdges() {
		F, T := d.FindSet(e)
		if F.Equal(T) {
			return true
		}
		d = d.Union(F, T)
	}
	return false
}

// DisjointSet is a type consist of VerSets,
// each of the VerSet is disjoint.
type DisjointSet []VerSet

func (g *graph) NewDisjointSet() (DisjointSet, error) {
	set := make([]VerSet, g.verticesCount)

	i := 0
	for _, v := range g.GetVertices() {
		s := NewVerSet()
		if err := s.Add(v); err != nil {
			return nil, err
		}
		set[i] = s
		i++
	}

	return set, nil
}

// Equal checks whether d is same as other or not.
func (d DisjointSet) Equal(other DisjointSet) bool {
	sort.Slice(d, func(i, j int) bool {
		var maxi, maxj Vertex
		for k := range d[i] {
			if k > maxi {
				maxi = k
			}
		}
		for k := range d[j] {
			if k > maxj {
				maxj = k
			}
		}
		return maxi < maxj
	})

	sort.Slice(other, func(i, j int) bool {
		var maxi, maxj Vertex
		for k := range other[i] {
			if k > maxi {
				maxi = k
			}
		}
		for k := range other[j] {
			if k > maxj {
				maxj = k
			}
		}
		return maxi < maxj
	})

	return reflect.DeepEqual(d, other)
}

// FindSet returns VerSets.
// Each VerSet contains the Vertex represented by `e.From` and `e.To`.
func (d DisjointSet) FindSet(e Edge) (F, T VerSet) {
	for _, set := range d {
		if _, ok := set[e.From]; ok {
			F = set
		}
		if _, ok := set[e.To]; ok {
			T = set
		}
	}
	return F, T
}

// Union unions 2 VerSets to one.
func (d DisjointSet) Union(A, B VerSet) DisjointSet {
	idxA, idxB := -1, -1
	for i, set := range d {
		if reflect.DeepEqual(set, A) {
			idxA = i
		}
		if reflect.DeepEqual(set, B) {
			idxB = i
		}

		if idxA >= 0 && idxB >= 0 {
			break
		}
	}

	if idxA > idxB {
		idxA, idxB = idxB, idxA
	}

	AB := make(VerSet, len(d[idxA])+len(d[idxB]))
	i := 0
	for v := range d[idxA] {
		AB[v] = struct{}{}
		i++
	}
	for v := range d[idxB] {
		AB[v] = struct{}{}
		i++
	}

	copy(d[idxB:], d[idxB+1:])
	d[len(d)-1] = nil
	d = d[:len(d)-1]

	copy(d[idxA:], d[idxA+1:])
	d[len(d)-1] = nil
	d = d[:len(d)-1]

	d = append(d, AB)

	return d
}
