package solutions

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
