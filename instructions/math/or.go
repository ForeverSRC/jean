package math

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type IOR struct {
	base.NoOperandsInstruction
}

func (i *IOR) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	res := v1 | v2
	stack.PushInt(res)
}

type LOR struct {
	base.NoOperandsInstruction
}

func (l *LOR) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	res := v1 | v2
	stack.PushLong(res)
}

func init() {
	ior := &IOR{}
	lor := &LOR{}

	factory.Factory.AddInstruction(0x80, func() base.Instruction {
		return ior
	})

	factory.Factory.AddInstruction(0x81, func() base.Instruction {
		return lor
	})
}
