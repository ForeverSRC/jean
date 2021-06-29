package loads

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type LLOAD struct {
	base.Index8Instruction
}

func (l *LLOAD) Execute(frame *jvmstack.Frame) {
	_lload(frame, l.Index)
}

func _lload(frame *jvmstack.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

// LLOAD_idx idx表示在局部变量表中的索引
type LLOAD_0 struct {
	base.NoOperandsInstruction
}

func (l *LLOAD_0) Execute(frame *jvmstack.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct {
	base.NoOperandsInstruction
}

func (l *LLOAD_1) Execute(frame *jvmstack.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct {
	base.NoOperandsInstruction
}

func (l *LLOAD_2) Execute(frame *jvmstack.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func (l *LLOAD_3) Execute(frame *jvmstack.Frame) {
	_lload(frame, 3)
}

func init() {
	factory.Factory.AddInstruction(0x16, func() base.Instruction {
		return &LLOAD{}
	})

	lload_0 := &LLOAD_0{}
	lload_1 := &LLOAD_1{}
	lload_2 := &LLOAD_2{}
	lload_3 := &LLOAD_3{}

	factory.Factory.AddInstruction(0x1e, func() base.Instruction {
		return lload_0
	})

	factory.Factory.AddInstruction(0x1f, func() base.Instruction {
		return lload_1
	})

	factory.Factory.AddInstruction(0x20, func() base.Instruction {
		return lload_2
	})

	factory.Factory.AddInstruction(0x21, func() base.Instruction {
		return lload_3
	})
}
