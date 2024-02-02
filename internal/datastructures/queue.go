package datastructures

type Queue []any

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) enqueue(val any) {
	*q = append(*q, val)
}

func (q *Queue) dequeue() any {
	if q.isEmpty() {
		return nil
	}

	val := (*q)[0]

	*q = (*q)[1:]

	return val
}
