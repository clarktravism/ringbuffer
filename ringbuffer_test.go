package ringbuffer

import "testing"

func TestNewRingBuffer(t *testing.T) {
	rb := NewRingBuffer[int](8)
	if rb.Cap() != 8 {
		t.Fail()
		t.Log("expected Cap() of 8")
	}
	if rb.Len() != 0 {
		t.Fail()
		t.Log("expected Len() of 0")
	}

	n := rb.Next() //returns zero-value of empty buffer
	if n != 0 {
		t.Fail()
		t.Log("expected n==0")
	}

	rb.Append(1)
	rb.Append(2)

	n = rb.Next()
	if n != 1 {
		t.Fail()
		t.Log("expected n==1")
	}

	n = rb.Next()
	if n != 2 {
		t.Fail()
		t.Log("expected n==2")
	}

	if !rb.Empty() {
		t.Fail()
		t.Log("expected Empty()==true")
	}

	if rb.Cap() != 8 {
		t.Fail()
		t.Log("expected Cap() of 8")
	}
	if rb.Len() != 0 {
		t.Fail()
		t.Log("expected Len() of 0")
	}
}

func TestAutoResize(t *testing.T) {
	rb := RingBuffer[int]{}
	rb.Append(1)
	rb.Append(2)
	rb.Append(3)

	if rb.Cap() != 3 {
		t.Fail()
		t.Log("expected Cap() of 3")
	}
	if rb.Len() != 3 {
		t.Fail()
		t.Log("expected Len() of 3")
	}

	n := rb.Next()
	if n != 1 {
		t.Fail()
		t.Log("expected n==1")
	}

	n = rb.Next()
	if n != 2 {
		t.Fail()
		t.Log("expected n==2")
	}

	n = rb.Next()
	if n != 3 {
		t.Fail()
		t.Log("expected n==3")
	}

	if !rb.Empty() {
		t.Fail()
		t.Log("expected Empty()==true")
	}

	if rb.Cap() != 3 {
		t.Fail()
		t.Log("expected Cap() of 3")
	}
	if rb.Len() != 0 {
		t.Fail()
		t.Log("expected Len() of 0")
	}
}

func TestLoopAround(t *testing.T) {

	rb := RingBuffer[int]{}

	rb.Append(1)
	if rb.Len() != 1 {
		t.Fail()
		t.Log("expected Len() of 1")
	}

	n := rb.Next()
	if n != 1 {
		t.Fail()
		t.Log("expected n==1")
	}
	if !rb.Empty() {
		t.Fail()
		t.Log("expected Empty()==true")
	}

	rb.Append(2)
	if rb.Len() != 1 {
		t.Fail()
		t.Log("expected Len() of 1")
	}

	n = rb.Next()
	if n != 2 {
		t.Fail()
		t.Log("expected n==2")
	}
	if !rb.Empty() {
		t.Fail()
		t.Log("expected Empty()==true")
	}

	rb.Append(3)
	if rb.Len() != 1 {
		t.Fail()
		t.Log("expected Len() of 1")
	}

	n = rb.Next()
	if n != 3 {
		t.Fail()
		t.Log("expected n==3")
	}
	if !rb.Empty() {
		t.Fail()
		t.Log("expected Empty()==true")
	}

	rb.Append(4)
	n = rb.Next()
	if n != 4 {
		t.Fail()
		t.Log("expected n==4")
	}
	if !rb.Empty() {
		t.Fail()
		t.Log("expected Empty()==true")
	}

	if rb.Cap() != 1 {
		t.Fail()
		t.Log("expected Cap() of 1")
	}
	if rb.Len() != 0 {
		t.Fail()
		t.Log("expected Len() of 0")
	}
}

func TestCopy(t *testing.T) {
	rb := NewRingBuffer[int](4)
	slice := []int{0, 0, 0, 0}

	n := rb.Copy(slice)
	if n != 0 {
		t.Fail()
		t.Log("expected n==0")
	}

	rb.Append(1)
	n = rb.Copy(slice)
	if n != 1 {
		t.Fail()
		t.Log("expected n==1")
	}
	if slice[0] != 1 {
		t.Fail()
		t.Log("expected slice[0]==1")
	}

	//loop the ring around
	rb.Next()
	rb.Append(2)
	rb.Next()
	rb.Append(3)
	rb.Next()
	rb.Append(4)
	rb.Append(5)

	n = rb.Copy(slice)
	if n != 2 {
		t.Fail()
		t.Log("expected n==2")
	}
	if slice[0] != 4 {
		t.Fail()
		t.Log("expected slice[0]==4")
	}
	if slice[1] != 5 {
		t.Fail()
		t.Log("expected slice[1]==5")
	}

	rb.Clear()
	n = rb.Copy(slice)
	if n != 0 {
		t.Fail()
		t.Log("expected n==0")
	}
}

func TestForNext(t *testing.T) {

	rb := NewRingBuffer[int](3)
	rb.Append(1)
	rb.Append(2)
	rb.Append(3)
	slice := make([]int, 0, 3)

	for !rb.Empty() {
		value := rb.Next()
		slice = append(slice, value)
	}

	if slice[0] != 1 || slice[1] != 2 || slice[2] != 3 || len(slice) != 3 {
		t.Fail()
		t.Log("expected slice[0]==1 slice[1]==2 slice[2]==3 len(slice)==3")
	}

	rb.Append(4)
	rb.Append(5)
	rb.Append(6)
	slice = slice[:0]

	for !rb.Empty() {
		value := rb.Next()
		slice = append(slice, value)
	}

	if slice[0] != 4 || slice[1] != 5 || slice[2] != 6 || len(slice) != 3 {
		t.Fail()
		t.Log("expected slice[0]==4 slice[1]==5 slice[2]==6 len(slice)==3")
	}
}
