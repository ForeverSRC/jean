package loads

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type ILOAD struct {
	base.Index8Instruction
}

func (i *ILOAD) Execute(frame *jvmstack.Frame) {
	_iload(frame, i.Index)
}

func _iload(frame *jvmstack.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

// ILOAD_idx idx表示在局部变量表中的索引

type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_0) Execute(frame *jvmstack.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_1) Execute(frame *jvmstack.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_2) Execute(frame *jvmstack.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (i *ILOAD_3) Execute(frame *jvmstack.Frame) {
	_iload(frame, 3)
}

func init() {
	factory.Factory.AddInstruction(0x15, func() base.Instruction {
		return &ILOAD{}
	})

	iload_0 := &ILOAD_0{}
	iload_1 := &ILOAD_1{}
	iload_2 := &ILOAD_2{}
	iload_3 := &ILOAD_3{}

	factory.Factory.AddInstruction(0x1a, func() base.Instruction {
		return iload_0
	})

	factory.Factory.AddInstruction(0x1b, func() base.Instruction {
		return iload_1
	})

	factory.Factory.AddInstruction(0x1c, func() base.Instruction {
		return iload_2
	})

	factory.Factory.AddInstruction(0x1d, func() base.Instruction {
		return iload_3
	})
}
