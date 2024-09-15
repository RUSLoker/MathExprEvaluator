package evaluator

// Stack defines a generic stack data structure
type Stack[T any] struct {
	items []T
}

// Push adds an item to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the item from the top of the stack
func (s *Stack[T]) Pop() (*T, bool) {
	if len(s.items) == 0 {
		return nil, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return &item, true
}

// PopN removes and returns up to the n items from the top of the stack
func (s *Stack[T]) PopN(n int) []T {
	n = min(n, len(s.items))
	items := s.items[len(s.items)-n:]
	s.items = s.items[:len(s.items)-n]
	return items
}

// Peek returns the item at the top of the stack without removing it
func (s *Stack[T]) Peek() (*T, bool) {
	if len(s.items) == 0 {
		return nil, false
	}
	return &s.items[len(s.items)-1], true
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}
