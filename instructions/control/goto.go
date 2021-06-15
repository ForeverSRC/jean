package control

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (g *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, g.Offset)
}

func init() {
	factory.Factory.AddInstruction(0xa7, func() base.Instruction {
		return &GOTO{}
	})
}
