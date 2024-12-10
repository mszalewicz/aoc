package main

type Queue[T any] struct {
	Elements []T
}

func (q *Queue[T]) Enqueue(value T) {
	q.Elements = append(q.Elements, value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.Elements) == 0 {
		var zero T
		return zero, false
	}
	value := q.Elements[0]
	q.Elements = q.Elements[1:]
	return value, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.Elements)
}