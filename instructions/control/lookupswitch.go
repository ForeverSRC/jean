package control

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	nparis        int32
	matchOffsets  []int32
}

func (ls *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	ls.defaultOffset = reader.ReadInt32()
	ls.nparis = reader.ReadInt32()
	ls.matchOffsets = reader.ReadInt32s(ls.nparis * 2)
}

func (ls *LOOKUP_SWITCH) Execute(frame *jvmstack.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < ls.nparis*2; i += 2 {
		if ls.matchOffsets[i] == key {
			offset := ls.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(ls.defaultOffset))
}

func init() {
	factory.Factory.AddInstruction(0xab, func() base.Instruction {
		return &LOOKUP_SWITCH{}
	})
}
