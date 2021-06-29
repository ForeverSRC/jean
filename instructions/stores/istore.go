package stores

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type ISTORE struct {
	base.Index8Instruction
}

func (i *ISTORE) Execute(frame *jvmstack.Frame) {
	_istore(frame, i.Index)
}

func _istore(frame *jvmstack.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

// ISTORE_idx idx表示在局部变量表中的索引

type ISTORE_0 struct {
	base.NoOperandsInstruction
}

func (i *ISTORE_0) Execute(frame *jvmstack.Frame) {
	_istore(frame, 0)
}

type ISTORE_1 struct {
	base.NoOperandsInstruction
}

func (i *ISTORE_1) Execute(frame *jvmstack.Frame) {
	_istore(frame, 1)
}

type ISTORE_2 struct {
	base.NoOperandsInstruction
}

func (i *ISTORE_2) Execute(frame *jvmstack.Frame) {
	_istore(frame, 2)
}

type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (i *ISTORE_3) Execute(frame *jvmstack.Frame) {
	_istore(frame, 3)
}

func init() {
	factory.Factory.AddInstruction(0x36, func() base.Instruction {
		return &ISTORE{}
	})

	istore_0 := &ISTORE_0{}
	istore_1 := &ISTORE_1{}
	istore_2 := &ISTORE_2{}
	istore_3 := &ISTORE_3{}

	factory.Factory.AddInstruction(0x3b, func() base.Instruction {
		return istore_0
	})

	factory.Factory.AddInstruction(0x3c, func() base.Instruction {
		return istore_1
	})

	factory.Factory.AddInstruction(0x3d, func() base.Instruction {
		return istore_2
	})

	factory.Factory.AddInstruction(0x3e, func() base.Instruction {
		return istore_3
	})
}
