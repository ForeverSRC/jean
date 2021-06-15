package loads

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type FLOAD struct {
	base.Index8Instruction
}

func (f *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, f.Index)
}

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

// FLOAD_idx idx表示在局部变量表中的索引
type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (f *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (f *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (f *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (f *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}

func init() {
	factory.Factory.AddInstruction(0x17, func() base.Instruction {
		return &FLOAD{}
	})

	fload_0 := &FLOAD_0{}
	fload_1 := &FLOAD_1{}
	fload_2 := &FLOAD_2{}
	fload_3 := &FLOAD_3{}

	factory.Factory.AddInstruction(0x22, func() base.Instruction {
		return fload_0
	})

	factory.Factory.AddInstruction(0x23, func() base.Instruction {
		return fload_1
	})

	factory.Factory.AddInstruction(0x24, func() base.Instruction {
		return fload_2
	})

	factory.Factory.AddInstruction(0x25, func() base.Instruction {
		return fload_3
	})
}
