package set

// Set is an interface should be implemented as a mathematically set.
type Set interface {
	// Add adds an `i` to the Set.
	Add(i interface{}) error

	// Remove removes an `i` from the Set.
	Remove(i interface{}) error

	// Cardinality returns the number of elements in the Set.
	Cardinality() int

	// Contains checks whether an in exists in the Set or not.
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
