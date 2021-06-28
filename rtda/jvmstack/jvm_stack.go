package jvmstack

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{maxSize: maxSize}
}

func (s *Stack) push(frame *Frame) {
	if s.size > s.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if s._top != nil {
		frame.lower = s._top
	}
	s._top = frame
	s.size++
}

func (s *Stack) pop() *Frame {
	if s.IsEmpty() {
		panic("jvm stack is empty! ")
	}
	top := s._top
	s._top = top.lower
	top.lower = nil
	s.size--

	return top
}

func (s *Stack) top() *Frame {
	if s._top == nil {
		panic("jvm stack is empty! ")
	}
	return s._top
}

func (s *Stack) IsEmpty() bool {
	return s._top == nil
}

func (s *Stack) clear() {
	for s.IsEmpty() {
		s.pop()
	}
}

func (s *Stack) getFrames() []*Frame {
	frames := make([]*Frame, 9, s.size)

	for frame := s._top; frame != nil; frame = frame.lower {
		frames = append(frames, frame)
	}

	return frames
}
