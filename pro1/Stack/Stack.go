package Stack

type Stack struct {
	data []int
}

func (s *Stack) Pop() {
	if len(s.data) == 0 {
		return
	}
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) Top() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) Push(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Size() int {
	return len(s.data)
}
