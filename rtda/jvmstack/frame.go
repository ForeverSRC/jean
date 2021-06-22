package jvmstack

import (
	"jean/rtda/heap"
)

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	method       *heap.Method
	thread       *Thread

	// the next instruction after the call
	nextPC int
}

func newFrame(t *Thread, method *heap.Method) *Frame {
	return &Frame{
		localVars:    NewLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
		thread:       t,
		method:       method,
	}
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) Method() *heap.Method {
	return f.method
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) NextPC() int {
	return f.nextPC
}

func (f *Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}

// RevertNextPC 让下一条指令重新指向当前指令
func (f *Frame) RevertNextPC() {
	f.nextPC = f.thread.pc
}
