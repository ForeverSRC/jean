package stores

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type FSTORE struct {
	base.Index8Instruction
}

func (f *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, f.Index)
}

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

// FSTORE_idx idx表示在局部变量表中的索引
type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (f *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (f *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (f *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (f *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}

func init() {
	factory.Factory.AddInstruction(0x38, func() base.Instruction {
		return &FSTORE{}
	})

	fstore_0 := &FSTORE_0{}
	fstore_1 := &FSTORE_1{}
	fstore_2 := &FSTORE_2{}
	fstore_3 := &FSTORE_3{}

	factory.Factory.AddInstruction(0x43, func() base.Instruction {
		return fstore_0
	})

	factory.Factory.AddInstruction(0x44, func() base.Instruction {
		return fstore_1
	})

	factory.Factory.AddInstruction(0x45, func() base.Instruction {
		return fstore_2
	})

	factory.Factory.AddInstruction(0x46, func() base.Instruction {
		return fstore_3
	})
}
