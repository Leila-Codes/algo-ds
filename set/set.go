package set

type Set[T comparable] map[T]struct{}

// New constructs a new set with the (optional) specified values contained within.
func New[T comparable](initValues ...T) Set[T] {
	set := Set[T]{}

	set.Add(initValues...)

	return set
}

// Contains returns whether the given item exists in the set or not.
func (s Set[T]) Contains(value T) bool {
	_, exists := (s)[value]
	return exists
}

// Add adds the provided value(s) to the set.
func (s Set[T]) Add(values ...T) {
	for _, v := range values {
		if _, exists := s[v]; !exists {
			s[v] = struct{}{}
		}
	}
}

// Remove removes the specified value(s) from the set.
func (s Set[T]) Remove(values ...T) {
	for _, v := range values {
		if _, exists := s[v]; exists {
			delete(s, v)
		}
	}
}
