package graph

import (
	"errors"
	"reflect"
)

// func (g *graph) ExistsCycle() bool {
//
// 	return false
// }

// func (g *graph) DisjointSet() bool {
//
// 	return false
// }

// Set is a set consists of the vertex or vertices
type Set map[Vertex]struct{}

// DisjointSet is an interface to execute Disjoint-Set Algorithm
type DisjointSet interface {
	MakeSet(g *graph) []Set
	Union(A, B Set) Set
	FindSet(e Edge, U []Set) (Set, Set)
}

// MakeSet returns a set of the vertices
// to solve by disjoint-set algorithm(initiallize)
func MakeSet(g *graph) []Set {
	U := make([]Set, g.verticesCount)

	for i, v := range g.GetVertices() {
		U[i] = Set{
			v: struct{}{},
		}
	}
	return U
}

// Union unions two sets into a set.
// AddVertex guarantees all vertices is different.
// func Union(A, B Set, U []Set) (Set, []Set) {
// 	C := make(Set, len(A)+len(B))
//
// 	for a := range A {
// 		C[a] = struct{}{}
// 	}
// 	for b := range B {
// 		C[b] = struct{}{}
// 	}
//
// 	return C, U
// 	// TODO: 結合する前の全体集合からAとBを消してUを追加する必要がある
// }

// Contains checks whether U contains A or not
func Contains(A Set, U []Set) bool {
	for _, u := range U {
		if reflect.DeepEqual(u, A) {
			return true
		}
	}
	return false
}

// Add adds A to U
func Add(A Set, U []Set) ([]Set, error) {
	if Contains(A, U) {
		return nil, errors.New("A is aleady in U")
	}

	return append(U, A), nil
}

// Remove removes A from U
func Remove(A Set, U []Set) ([]Set, error) {
	if !Contains(A, U) {
		return nil, errors.New("A does not exist in U")
	}

	i := 0
	for j, u := range U {
		if reflect.DeepEqual(u, A) {
			i = j
			break
		}
	}

	copy(U[i:], U[i+1:])
	U[len(U)-1] = nil
	U = U[:len(U)-1]

	return U, nil
}

// FindSet returns the sets contains the vertex
// both ends of the edge `e`
func FindSet(e Edge, U []Set) (A, B Set) {
	for _, set := range U {
		if _, ok := set[e.From]; ok {
			A = set
		}
		if _, ok := set[e.To]; ok {
			B = set
		}
	}

	return A, B
}
