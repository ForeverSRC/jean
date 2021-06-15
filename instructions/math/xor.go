package math

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type IXOR struct {
	base.NoOperandsInstruction
}

func (i *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	res := v1 ^ v2
	stack.PushInt(res)
}

type LXOR struct {
	base.NoOperandsInstruction
}

func (l *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	res := v1 ^ v2
	stack.PushLong(res)
}

func init() {
	ixor := &IXOR{}
	lxor := &LXOR{}

	factory.Factory.AddInstruction(0x82, func() base.Instruction {
		return ixor
	})

	factory.Factory.AddInstruction(0x83, func() base.Instruction {
		return lxor
	})
}
