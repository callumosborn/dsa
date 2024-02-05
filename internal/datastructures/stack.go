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

func (s *Stack) Iterator() Interator {
	return &StackIterator{
		index: len(*s) - 1,
		stack: *s,
	}
}

type StackIterator struct {
	index int
	stack Stack
}

func (si *StackIterator) HasNext() bool {
	return si.index >= 0
}

func (si *StackIterator) GetNext() any {
	if !si.HasNext() {
		return nil
	}

	val := si.stack[si.index]
	si.index--

	return val
}
