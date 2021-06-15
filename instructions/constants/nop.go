package constants

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (nop *NOP) Execute(frame *rtda.Frame) {
	//do nothing
}

func init() {
	nop := &NOP{}
	factory.Factory.AddInstruction(0x00, func() base.Instruction {
		return nop
	})
}
