package datastructures

type Stack []any

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(value any) {
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
