package ringbuffer

//RingBuffer data structure that will reuse elements that were already read, and will auto-expand when the length of unread elements exceeds the buffer.
//Ready to use as RingBuffer[T]{} but it recommended to set a capacity with NewRingBuffer().
type RingBuffer[T any] struct {
	buffer     []T
	head, tail int
}

//NewRingBuffer will create a RingBuffer with the capacity provided.
func NewRingBuffer[T any](cap int) RingBuffer[T] {
	return RingBuffer[T]{buffer: make([]T, cap+1)}
}

func (rb RingBuffer[T]) advance(i int) int {
	if i < len(rb.buffer)-1 {
		return i + 1
	}
	return 0
}

//Append adds a new value to the buffer.
func (rb *RingBuffer[T]) Append(value T) {
	if len(rb.buffer) == 0 {
		rb.buffer = make([]T, 2)
	}
	ntail := rb.advance(rb.tail)

	//resize the buffer
	if ntail == rb.head {
		nBuffer := make([]T, len(rb.buffer)*2)
		rb.tail = rb.Copy(nBuffer)
		rb.head = 0
		rb.buffer = nBuffer
		ntail = rb.tail + 1
	}

	rb.buffer[rb.tail] = value
	rb.tail = ntail
}

//Copy copies the contents of the RingBuffer to a slice of the same type. Copy returns the number of elements copied.
func (rb RingBuffer[T]) Copy(slice []T) int {
	if rb.head == rb.tail {
		return 0
	}
	if rb.head < rb.tail {
		return copy(slice, rb.buffer[rb.head:rb.tail])
	}
	n := copy(slice, rb.buffer[rb.head:])
	return n + copy(slice[n:], rb.buffer[0:rb.tail])
}

//Next returns the next value in the buffer and advances the buffer if it can.
//If the buffer is empty, it will return the zero value of the RingBuffer type.
func (rb *RingBuffer[T]) Next() T {
	if rb.head == rb.tail {
		return *new(T)
	}
	value := rb.buffer[rb.head]
	rb.head = rb.advance(rb.head)
	return value
}

//Empty returns if the buffer is empty.
func (rb RingBuffer[T]) Empty() bool {
	return rb.head == rb.tail
}

//Len returns the current size of the buffer.
func (rb RingBuffer[T]) Len() int {
	if rb.head <= rb.tail {
		return rb.tail - rb.head
	}
	return len(rb.buffer) - rb.head + rb.tail
}

//Len returns the current capacity of the buffer.
func (rb RingBuffer[T]) Cap() int {
	return len(rb.buffer) - 1
}

//Clear resets the buffer.
func (rb *RingBuffer[T]) Clear() {
	rb.head, rb.tail = 0, 0
}
