package loads

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type DLOAD struct {
	base.Index8Instruction
}

func (d *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, d.Index)
}

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

// DLOAD_idx idx表示在局部变量表中的索引
type DLOAD_0 struct {
	base.NoOperandsInstruction
}

func (d *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

func (d *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct {
	base.NoOperandsInstruction
}

func (d *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (d *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}

func init() {
	factory.Factory.AddInstruction(0x18, func() base.Instruction {
		return &DLOAD{}
	})

	dload_0 := &DLOAD_0{}
	dload_1 := &DLOAD_1{}
	dload_2 := &DLOAD_2{}
	dload_3 := &DLOAD_3{}

	factory.Factory.AddInstruction(0x26, func() base.Instruction {
		return dload_0
	})

	factory.Factory.AddInstruction(0x27, func() base.Instruction {
		return dload_1
	})

	factory.Factory.AddInstruction(0x28, func() base.Instruction {
		return dload_2
	})

	factory.Factory.AddInstruction(0x29, func() base.Instruction {
		return dload_3
	})
}
