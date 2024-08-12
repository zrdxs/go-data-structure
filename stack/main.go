package main

// Doing with array
type Stack struct {
	values    []int
	maxLenght int
	top       int
}

func (s *Stack) Push(value int) {
	if len(s.values) >= s.maxLenght {
		return
	} else {
		s.top++
		s.values[s.top] = value
		return
	}
}

func (s *Stack) Pop() (value int) {
	if len(s.values) <= 0 {
		return 0
	} else {
		value = s.values[s.top]
		s.top = s.top - 1
		return value
	}
}

func (s *Stack) Peek() (value int) {
	return s.values[s.top]
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) <= 0
}

func main() {

}
