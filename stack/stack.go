package stack

type Stack []int

// Push method to add an element to the stack
func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

// Pop method to remove and return the top element from the stack
func (s *Stack) Pop() (int, bool) {
	if len(*s) == 0 {
		return 0, false // return zero value and false if stack is empty
	}

	// Get the top element
	index := len(*s) - 1
	element := (*s)[index]

	// Remove the top element
	*s = (*s)[:index]
	return element, true
}

// Peek method to return the top element without removing it
func (s *Stack) Peek() (int, bool) {
	if len(*s) == 0 {
		return 0, false // return zero value and false if stack is empty
	}

	// Get the top element
	index := len(*s) - 1
	element := (*s)[index]
	return element, true
}

func (s *Stack) length() int {
	return len(*s)
}
