package collections

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: make([]T, 0, 1000)}
}

func (queue *Queue[T]) Put(t T) {
	queue.data = append(queue.data, t)
}

func (queue *Queue[T]) Get() (T, bool) {
	if len(queue.data) == 0 {
		var zero T
		return zero, false
	}

	last := queue.data[0]
	queue.data = queue.data[1:]

	return last, true
}

func (queue *Queue[T]) Drain() []T {
	drain := make([]T, len(queue.data))
	idx := 0
	for {
		val, ok := queue.Get()

		if !ok {
			break
		}

		drain[idx] = val
		idx++
	}

	return drain
}
