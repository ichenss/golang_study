package Queue

type Queue struct {
	data []int
}

func (q *Queue) Push(val int) {
	q.data = append(q.data, val)
}

func (q *Queue) Pop() int {
	if len(q.data) == 0 {
		return -1
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func (q *Queue) Front() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0]
}

func (q *Queue) Back() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[len(q.data)-1]
}

func (q *Queue) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue) Size() int {
	return len(q.data)
}
