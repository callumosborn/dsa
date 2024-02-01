package datastructures

// Stack is a user-defined type that has the underlying type of
// a slice of the empty interface literal.
type Stack []any

// isEmpty checks if the stack is empty.
// It returns true if the stack is empty, otherwise returns false.
// The method is associated with the Stack type and operates on a pointer to a Stack.
func (s *Stack) isEmpty() bool {
	// The * is the indirection operator.
	// It precedes a variable of pointer type and returns the pointed-to value.
	// This is called dereferencing.
	//
	// s is a variable of type pointer-to-Stack.
	// s is dereferenced using *s, the pointed to value being the slice is returned.
	//
	// The length of the slice is checked.
	return len(*s) == 0
}

func (s *Stack) push(value any) {
	// append add elements (in this case: value) to the end of a slice.
	//
	// s is dereferenced using *s.
	//
	// value is appended to the slice s.
	// s is being assigned the new slice.
	*s = append(*s, value)
}

func (s *Stack) pop() any {
	if s.isEmpty() {
		return nil
	}

	val := (*s)[len(*s)-1:]

	*s = (*s)[:len(*s)-1]

	return val[0]
}
