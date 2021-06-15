package extended

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type GOTO_W struct {
	offset int
}

func (g *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	g.offset = int(reader.ReadInt32())
}

func (g *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, g.offset)
}

func init() {
	factory.Factory.AddInstruction(0xc8, func() base.Instruction {
		return &GOTO_W{}
	})
}
