package comparisons

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type IFEQ struct {
	base.BranchInstruction
}

func (ifcond *IFEQ) Execute(frame *rtda.Frame) {
	_ifcond(frame, func(i int32) bool {
		return i == 0
	}, ifcond.Offset)
}

type IFNE struct {
	base.BranchInstruction
}

func (ifcond *IFNE) Execute(frame *rtda.Frame) {
	_ifcond(frame, func(i int32) bool {
		return i != 0
	}, ifcond.Offset)
}

type IFLT struct {
	base.BranchInstruction
}

func (ifcond *IFLT) Execute(frame *rtda.Frame) {
	_ifcond(frame, func(i int32) bool {
		return i < 0
	}, ifcond.Offset)
}

type IFLE struct {
	base.BranchInstruction
}

func (ifcond *IFLE) Execute(frame *rtda.Frame) {
	_ifcond(frame, func(i int32) bool {
		return i <= 0
	}, ifcond.Offset)
}

type IFGT struct {
	base.BranchInstruction
}

func (ifcond *IFGT) Execute(frame *rtda.Frame) {
	_ifcond(frame, func(i int32) bool {
		return i > 0
	}, ifcond.Offset)
}

type IFGE struct {
	base.BranchInstruction
}

func (ifcond *IFGE) Execute(frame *rtda.Frame) {
	_ifcond(frame, func(i int32) bool {
		return i >= 0
	}, ifcond.Offset)
}

func _ifcond(frame *rtda.Frame, cond func(int32) bool, offset int) {
	val := frame.OperandStack().PopInt()
	if cond(val) {
		base.Branch(frame, offset)
	}
}

func init() {
	factory.Factory.AddInstruction(0x99, func() base.Instruction {
		return &IFEQ{}
	})

	factory.Factory.AddInstruction(0x9a, func() base.Instruction {
		return &IFNE{}
	})

	factory.Factory.AddInstruction(0x9b, func() base.Instruction {
		return &IFLT{}
	})

	factory.Factory.AddInstruction(0x9c, func() base.Instruction {
		return &IFGE{}
	})

	factory.Factory.AddInstruction(0x9d, func() base.Instruction {
		return &IFGT{}
	})

	factory.Factory.AddInstruction(0x9e, func() base.Instruction {
		return &IFLE{}
	})
}
