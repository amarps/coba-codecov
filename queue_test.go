package queue

import "testing"

func TestPush(t *testing.T) {
	q := NewQueue()
	q.Push(5)
	if q.Len() != 1 {
		t.Errorf("Expected %d, got: %d", 1, q.Len())
	}
}

func Test(t *testing.T) {
	q := NewQueue()
	if !q.IsEmpty() {
		t.Errorf("Expected %t, got: %t", true, q.IsEmpty())
	}
	q.Push(4)
	q.Empty()
	if !q.IsEmpty() {
		t.Errorf("Expected %t, got: %t", true, q.IsEmpty())
	}
	q.Push(3)
	res := q.Pop() == 3
	if !res {
		t.Errorf("Expected %t, got: %t", true, res)
	}
	if !q.IsEmpty() {
		t.Errorf("Expected %t, got: %t", true, q.IsEmpty())
	}
}
