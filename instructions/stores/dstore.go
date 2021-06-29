package stores

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type DSTORE struct {
	base.Index8Instruction
}

func (d *DSTORE) Execute(frame *jvmstack.Frame) {
	_dstore(frame, d.Index)
}

func _dstore(frame *jvmstack.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

// DSTORE_idx idx表示在局部变量表中的索引
type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (d *DSTORE_0) Execute(frame *jvmstack.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct {
	base.NoOperandsInstruction
}

func (d *DSTORE_1) Execute(frame *jvmstack.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (d *DSTORE_2) Execute(frame *jvmstack.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (d *DSTORE_3) Execute(frame *jvmstack.Frame) {
	_dstore(frame, 3)
}

func init() {
	factory.Factory.AddInstruction(0x39, func() base.Instruction {
		return &DSTORE{}
	})

	dstore_0 := &DSTORE_0{}
	dstore_1 := &DSTORE_1{}
	dstore_2 := &DSTORE_2{}
	dstore_3 := &DSTORE_3{}

	factory.Factory.AddInstruction(0x47, func() base.Instruction {
		return dstore_0
	})

	factory.Factory.AddInstruction(0x48, func() base.Instruction {
		return dstore_1
	})

	factory.Factory.AddInstruction(0x49, func() base.Instruction {
		return dstore_2
	})

	factory.Factory.AddInstruction(0x4a, func() base.Instruction {
		return dstore_3
	})
}
