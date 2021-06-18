package loads

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type ALOAD struct {
	base.Index8Instruction
}

func (a *ALOAD) Execute(frame *jvmstack.Frame) {
	_aload(frame, a.Index)
}

func _aload(frame *jvmstack.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

// ALOAD_idx idx表示在局部变量表中的索引
type ALOAD_0 struct {
	base.NoOperandsInstruction
}

func (a *ALOAD_0) Execute(frame *jvmstack.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct {
	base.NoOperandsInstruction
}

func (a *ALOAD_1) Execute(frame *jvmstack.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct {
	base.NoOperandsInstruction
}

func (a *ALOAD_2) Execute(frame *jvmstack.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (a *ALOAD_3) Execute(frame *jvmstack.Frame) {
	_aload(frame, 3)
}

func init() {
	factory.Factory.AddInstruction(0x19, func() base.Instruction {
		return &ALOAD{}
	})

	aload_0 := &ALOAD_0{}
	aload_1 := &ALOAD_1{}
	aload_2 := &ALOAD_2{}
	aload_3 := &ALOAD_3{}

	factory.Factory.AddInstruction(0x2a, func() base.Instruction {
		return aload_0
	})

	factory.Factory.AddInstruction(0x2b, func() base.Instruction {
		return aload_1
	})

	factory.Factory.AddInstruction(0x2c, func() base.Instruction {
		return aload_2
	})

	factory.Factory.AddInstruction(0x2d, func() base.Instruction {
		return aload_3
	})
}
