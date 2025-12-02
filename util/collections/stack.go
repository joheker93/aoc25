package collections

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0, 1000)}
}

func (stack *Stack[T]) Push(t T) {
	stack.data = append(stack.data, t)
}

func (stack *Stack[T]) Pop() (T, bool) {
	len := len(stack.data)
	if len == 0 {
		var zero T
		return zero, false
	}

	last := stack.data[len-1]
	stack.data = stack.data[:len-1]
	return last, true
}
