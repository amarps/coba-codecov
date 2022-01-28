package queue

import "testing"

func TestPush(t *testing.T) {
	q := NewQueue()
	q.Push(5)
	if q.Len() != 1 {
		t.Errorf("Expected %d, got: %d", 1, q.Len())
	}
}
