package comparisons

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type LCMP struct {
	base.NoOperandsInstruction
}

func (lcmp *LCMP) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}

func init() {
	lcmp := &LCMP{}
	factory.Factory.AddInstruction(0x94, func() base.Instruction {
		return lcmp
	})
}
