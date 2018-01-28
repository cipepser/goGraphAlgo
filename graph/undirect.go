package graph

// func (g *graph) ExistsCycle() bool {
//
// 	return false
// }

// func (g *graph) DisjointSet() bool {
//
// 	return false
// }

type DisjointSet interface {
	MakeSet() []map[Vertex]struct{}
	Union()
	FindSet()
}

// MakeSet returns a set of the vertices
// to solve by disjoint-set algorithm(initiallize)
func MakeSet(g *graph) []map[Vertex]struct{} {
	sets := make([]map[Vertex]struct{}, g.verticesCount)

	for i, v := range g.GetVertices() {
		sets[i] = map[Vertex]struct{}{
			v: struct{}{},
		}
	}
	return sets
}

// Union unions two sets into a set.
// AddVertex guarantees all vertices is different.
func Union(A, B []map[Vertex]struct{}) []map[Vertex]struct{} {
	U := make([]map[Vertex]struct{}, len(A)+len(B))

	i := 0
	for _, a := range A {
		U[i] = a
		i++
	}
	for _, b := range B {
		U[i] = b
		i++
	}

	return U
}

func FindSet() {

}
