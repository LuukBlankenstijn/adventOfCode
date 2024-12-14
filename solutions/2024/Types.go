package _024

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type State struct {
	location  Location
	direction Direction
}

type Location struct {
	y int
	x int
}

type Set[T comparable] map[T]struct{}

// Add adds an item to the set.
func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

// Remove removes an item from the set.
func (s Set[T]) Remove(item T) {
	delete(s, item)
}

// Contains checks if the set contains the given item.
func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

// Items returns all items in the set as a slice.
func (s Set[T]) Items() []T {
	items := make([]T, 0, len(s))
	for item := range s {
		items = append(items, item)
	}
	return items
}

func (s Set[T]) Copy() Set[T] {
	items := s.Items()
	newSet := make(Set[T])
	for _, item := range items {
		newSet.Add(item)
	}
	return newSet
}

func (s Set[T]) difference(subtract Set[T]) Set[T] {
	newSet := make(Set[T])
	for _, item := range s.Items() {
		if !subtract.Contains(item) {
			newSet.Add(item)
		}
	}
	return newSet
}

func (s Set[T]) Union(second Set[T]) {
	for item := range second {
		s.Add(item)
	}
}
