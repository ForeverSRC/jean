package stores

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type LSTORE struct {
	base.Index8Instruction
}

func (l *LSTORE) Execute(frame *jvmstack.Frame) {
	_lstore(frame, l.Index)
}

func _lstore(frame *jvmstack.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

// LSTORE_idx idx表示在局部变量表中的索引
type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (l *LSTORE_0) Execute(frame *jvmstack.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (l *LSTORE_1) Execute(frame *jvmstack.Frame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (l *LSTORE_2) Execute(frame *jvmstack.Frame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (l *LSTORE_3) Execute(frame *jvmstack.Frame) {
	_lstore(frame, 3)
}

func init() {
	factory.Factory.AddInstruction(0x37, func() base.Instruction {
		return &LSTORE{}
	})

	lstore_0 := &LSTORE_0{}
	lstore_1 := &LSTORE_1{}
	lstore_2 := &LSTORE_2{}
	lstore_3 := &LSTORE_3{}

	factory.Factory.AddInstruction(0x3f, func() base.Instruction {
		return lstore_0
	})

	factory.Factory.AddInstruction(0x40, func() base.Instruction {
		return lstore_1
	})

	factory.Factory.AddInstruction(0x41, func() base.Instruction {
		return lstore_2
	})

	factory.Factory.AddInstruction(0x42, func() base.Instruction {
		return lstore_3
	})
}
