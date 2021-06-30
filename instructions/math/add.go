package math

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type DADD struct {
	base.NoOperandsInstruction
}

func (d *DADD) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()

	v2 := stack.PopDouble()
	v1 := stack.PopDouble()

	res := v1 + v2

	stack.PushDouble(res)
}

type FADD struct {
	base.NoOperandsInstruction
}

func (f *FADD) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()

	v2 := stack.PopFloat()
	v1 := stack.PopFloat()

	res := v1 + v2

	stack.PushFloat(res)
}

type IADD struct {
	base.NoOperandsInstruction
}

func (i *IADD) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()

	v2 := stack.PopInt()
	v1 := stack.PopInt()

	res := v1 + v2

	stack.PushInt(res)
}

type LADD struct {
	base.NoOperandsInstruction
}

func (l *LADD) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()

	v2 := stack.PopLong()
	v1 := stack.PopLong()

	res := v1 + v2

	stack.PushLong(res)
}

func init() {
	iadd := &IADD{}
	ladd := &LADD{}
	fadd := &FADD{}
	dadd := &DADD{}

	factory.Factory.AddInstruction(0x60, func() base.Instruction {
		return iadd
	})

	factory.Factory.AddInstruction(0x61, func() base.Instruction {
		return ladd
	})

	factory.Factory.AddInstruction(0x62, func() base.Instruction {
		return fadd
	})

	factory.Factory.AddInstruction(0x63, func() base.Instruction {
		return dadd
	})
}
