package min_stack

// SolveMinStack supports push, pop, top, and getMin in O(1). O(n) space.
type SolveMinStack struct {
	stack    []int
	minStack []int
}

func NewSolveMinStack() SolveMinStack {
	return SolveMinStack{}
}

func (s *SolveMinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if len(s.minStack) == 0 || val <= s.minStack[len(s.minStack)-1] {
		s.minStack = append(s.minStack, val)
	}
}

func (s *SolveMinStack) Pop() {
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if top == s.minStack[len(s.minStack)-1] {
		s.minStack = s.minStack[:len(s.minStack)-1]
	}
}

func (s *SolveMinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *SolveMinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}
