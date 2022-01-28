package queue

type Queue struct {
	collections []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (Q *Queue) Empty() {
	Q.collections = []int{}
}

func (Q *Queue) IsEmpty() bool {
	return Q.Len() == 0
}

func (Q *Queue) Len() int {
	return len(Q.collections)
}

func (Q *Queue) Push(v int) {
	Q.collections = append(Q.collections, v)
}

func (Q *Queue) Pop(v int) int {
	res := Q.collections[0]
	Q.collections = Q.collections[1:]
	return res
}
