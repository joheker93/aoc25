package collections

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

func FromSlice[T comparable](ts []T) *Set[T] {
	set := NewSet[T]()

	for _, t := range ts {
		set.Add(t)
	}

	return set
}

func (set *Set[T]) Add(t T) bool {
	if set.Contains(t) {
		return false
	}

	set.data[t] = struct{}{}
	return true

}

func (set *Set[T]) Contains(t T) bool {
	_, exists := set.data[t]
	return exists
}

func (set *Set[T]) Remove(t T) bool {

	if !set.Contains(t) {
		return false
	}

	delete(set.data, t)
	return true
}

func (set *Set[T]) Size() int {
	return len(set.data)
}

func (set *Set[T]) Values() []T {
	keys := make([]T, 0, len(set.data))
	for k := range set.data {
		keys = append(keys, k)
	}
	return keys
}
