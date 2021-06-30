package math

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type IAND struct {
	base.NoOperandsInstruction
}

func (i *IAND) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	res := v1 & v2
	stack.PushInt(res)
}

type LAND struct {
	base.NoOperandsInstruction
}

func (l *LAND) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	res := v1 & v2
	stack.PushLong(res)
}

func init() {
	iand := &IAND{}
	land := &LAND{}

	factory.Factory.AddInstruction(0x7e, func() base.Instruction {
		return iand
	})

	factory.Factory.AddInstruction(0x7f, func() base.Instruction {
		return land
	})
}
