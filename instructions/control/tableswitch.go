package control

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (ts *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	ts.defaultOffset = reader.ReadInt32()
	ts.low = reader.ReadInt32()
	ts.high = reader.ReadInt32()
	jumpOffset := ts.high - ts.low + 1
	ts.jumpOffsets = reader.ReadInt32s(jumpOffset)
}

func (ts *TABLE_SWITCH) Execute(frame *jvmstack.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if ts.low <= index && index <= ts.high {
		offset = int(ts.jumpOffsets[index-ts.low])
	} else {
		offset = int(ts.defaultOffset)
	}

	base.Branch(frame, offset)
}

func init() {
	factory.Factory.AddInstruction(0xaa, func() base.Instruction {
		return &TABLE_SWITCH{}
	})
}
