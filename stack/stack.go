package stack

type Stack[T int | string] []T

// Push method to add an element to the stack
func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

// Pop method to remove and return the top element from the stack
func (s *Stack[T]) Pop() (item T, ok bool) {
	if len(*s) == 0 {
		return item, false // return zero value and false if stack is empty
	}

	// Get the top element
	index := len(*s) - 1
	item = (*s)[index]

	// Remove the top element
	*s = (*s)[:index]
	return item, true
}

// Peek method to return the top element without removing it
func (s *Stack[T]) Peek() (item T, ok bool) {
	if len(*s) == 0 {
		return item, false // return zero value and false if stack is empty
	}

	// Get the top element
	index := len(*s) - 1
	item = (*s)[index]
	return item, true
}

func (s *Stack[T]) length() int {
	return len(*s)
}
