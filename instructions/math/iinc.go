package math

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type IINC struct {
	Index uint
	Const int32
}

func (i *IINC) FetchOperands(reader *base.BytecodeReader) {
	i.Index = uint(reader.ReadUint8())
	i.Const = int32(reader.ReadInt8())
}

func (i *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(i.Index)
	val += i.Const
	localVars.SetInt(i.Index, val)
}

func init() {
	factory.Factory.AddInstruction(0x84, func() base.Instruction {
		return &IINC{}
	})
}
