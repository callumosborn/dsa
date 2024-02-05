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

func (q *Queue) Iterator() Interator {
	return &QueueIterator{
		queue: *q,
	}
}

type QueueIterator struct {
	index int
	queue Queue
}

func (qi *QueueIterator) HasNext() bool {
	return qi.index < len(qi.queue)
}

func (qi *QueueIterator) GetNext() any {
	if !qi.HasNext() {
		return nil
	}

	val := qi.queue[qi.index]

	qi.index++

	return val
}
