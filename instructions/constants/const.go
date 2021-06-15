package constants

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type ACONST_NULL struct {
	base.NoOperandsInstruction
}

func (bc *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

type DCONST_0 struct {
	base.NoOperandsInstruction
}

func (bc *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DCONST_1 struct {
	base.NoOperandsInstruction
}

func (bc *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

type FCONST_0 struct {
	base.NoOperandsInstruction
}

func (bc *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type FCONST_1 struct {
	base.NoOperandsInstruction
}

func (bc *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type FCONST_2 struct {
	base.NoOperandsInstruction
}

func (bc *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

type ICONST_M1 struct {
	base.NoOperandsInstruction
}

func (bc *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct {
	base.NoOperandsInstruction
}

func (bc *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

type ICONST_1 struct {
	base.NoOperandsInstruction
}

func (bc *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

type ICONST_2 struct {
	base.NoOperandsInstruction
}

func (bc *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

type ICONST_3 struct {
	base.NoOperandsInstruction
}

func (bc *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

type ICONST_4 struct {
	base.NoOperandsInstruction
}

func (bc *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

type ICONST_5 struct {
	base.NoOperandsInstruction
}

func (bc *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

type LCONST_0 struct {
	base.NoOperandsInstruction
}

func (bc *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (bc *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}

func init() {
	aconst_null := &ACONST_NULL{}
	iconst_m1 := &ICONST_M1{}
	iconst_0 := &ICONST_0{}
	iconst_1 := &ICONST_1{}
	iconst_2 := &ICONST_2{}
	iconst_3 := &ICONST_3{}
	iconst_4 := &ICONST_4{}
	iconst_5 := &ICONST_5{}
	lconst_0 := &LCONST_0{}
	lconst_1 := &LCONST_1{}
	fconst_0 := &FCONST_0{}
	fconst_1 := &FCONST_1{}
	fconst_2 := &FCONST_2{}
	dconst_0 := &DCONST_0{}
	dconst_1 := &DCONST_1{}
	factory.Factory.AddInstruction(0x01, func() base.Instruction {
		return aconst_null
	})

	factory.Factory.AddInstruction(0x02, func() base.Instruction {
		return iconst_m1
	})

	factory.Factory.AddInstruction(0x03, func() base.Instruction {
		return iconst_0
	})

	factory.Factory.AddInstruction(0x04, func() base.Instruction {
		return iconst_1
	})

	factory.Factory.AddInstruction(0x05, func() base.Instruction {
		return iconst_2
	})

	factory.Factory.AddInstruction(0x06, func() base.Instruction {
		return iconst_3
	})

	factory.Factory.AddInstruction(0x07, func() base.Instruction {
		return iconst_4
	})

	factory.Factory.AddInstruction(0x08, func() base.Instruction {
		return iconst_5
	})

	factory.Factory.AddInstruction(0x09, func() base.Instruction {
		return lconst_0
	})

	factory.Factory.AddInstruction(0x0a, func() base.Instruction {
		return lconst_1
	})

	factory.Factory.AddInstruction(0x0b, func() base.Instruction {
		return fconst_0
	})

	factory.Factory.AddInstruction(0x0c, func() base.Instruction {
		return fconst_1
	})

	factory.Factory.AddInstruction(0x0d, func() base.Instruction {
		return fconst_2
	})

	factory.Factory.AddInstruction(0x0e, func() base.Instruction {
		return dconst_0
	})

	factory.Factory.AddInstruction(0x0f, func() base.Instruction {
		return dconst_1
	})
}
