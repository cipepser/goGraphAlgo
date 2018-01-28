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
