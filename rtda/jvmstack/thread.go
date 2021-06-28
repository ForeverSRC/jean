package jvmstack

import "jean/rtda/heap"

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		// todo 命令行指定虚拟机栈大小
		stack: newStack(1024),
	}
}

func (t *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(t, method)
}

func (t *Thread) PushFrame(frame *Frame) {
	t.stack.push(frame)
}

func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top()
}

func (t *Thread) PC() int {
	return t.pc
}

func (t *Thread) SetPC(pc int) {
	t.pc = pc
}

func (t *Thread) TopFrame() *Frame {
	return t.CurrentFrame()
}

func (t *Thread) IsStackEmpty() bool {
	return t.stack.IsEmpty()
}

func (t *Thread) ClearStack() {
	t.stack.clear()
}

func (t *Thread) GetFrames() []*Frame {
	return t.stack.getFrames()
}
