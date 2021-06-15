package stores

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type ASTORE struct {
	base.Index8Instruction
}

func (a *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, a.Index)
}

func _astore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}

// ASTORE_idx idx表示在局部变量表中的索引
type ASTORE_0 struct {
	base.NoOperandsInstruction
}

func (a *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

type ASTORE_1 struct {
	base.NoOperandsInstruction
}

func (a *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

type ASTORE_2 struct {
	base.NoOperandsInstruction
}

func (a *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func (a *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}

func init() {
	factory.Factory.AddInstruction(0x3a, func() base.Instruction {
		return &ASTORE{}
	})

	astore_0 := &ASTORE_0{}
	astore_1 := &ASTORE_1{}
	astore_2 := &ASTORE_2{}
	astore_3 := &ASTORE_3{}

	factory.Factory.AddInstruction(0x4b, func() base.Instruction {
		return astore_0
	})

	factory.Factory.AddInstruction(0x4c, func() base.Instruction {
		return astore_1
	})

	factory.Factory.AddInstruction(0x4d, func() base.Instruction {
		return astore_2
	})

	factory.Factory.AddInstruction(0x4e, func() base.Instruction {
		return astore_3
	})
}
