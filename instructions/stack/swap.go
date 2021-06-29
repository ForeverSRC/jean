package stack

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type SWAP struct {
	base.NoOperandsInstruction
}

func (s *SWAP) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()

	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}

func init() {
	swap := &SWAP{}
	factory.Factory.AddInstruction(0x5f, func() base.Instruction {
		return swap
	})
}
