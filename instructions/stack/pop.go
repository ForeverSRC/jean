package stack

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type POP struct {
	base.NoOperandsInstruction
}

func (p *POP) Execute(frame *jvmstack.Frame) {
	frame.OperandStack().PopSlot()
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (p *POP2) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}

func init() {
	pop := &POP{}
	pop2 := &POP2{}

	factory.Factory.AddInstruction(0x57, func() base.Instruction {
		return pop
	})

	factory.Factory.AddInstruction(0x58, func() base.Instruction {
		return pop2
	})
}
