package math

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type DMUL struct {
	base.NoOperandsInstruction
}

func (d *DMUL) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()

	v2 := stack.PopDouble()
	v1 := stack.PopDouble()

	res := v1 * v2

	stack.PushDouble(res)
}

type FMUL struct {
	base.NoOperandsInstruction
}

func (f *FMUL) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()

	v2 := stack.PopFloat()
	v1 := stack.PopFloat()

	res := v1 * v2

	stack.PushFloat(res)
}

type IMUL struct {
	base.NoOperandsInstruction
}

func (i *IMUL) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()

	v2 := stack.PopInt()
	v1 := stack.PopInt()

	res := v1 * v2

	stack.PushInt(res)
}

type LMUL struct {
	base.NoOperandsInstruction
}

func (l *LMUL) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()

	v2 := stack.PopLong()
	v1 := stack.PopLong()

	res := v1 * v2

	stack.PushLong(res)
}

func init() {
	imul := &IMUL{}
	lmul := &LMUL{}
	fmul := &FMUL{}
	dmul := &DMUL{}

	factory.Factory.AddInstruction(0x68, func() base.Instruction {
		return imul
	})

	factory.Factory.AddInstruction(0x69, func() base.Instruction {
		return lmul
	})

	factory.Factory.AddInstruction(0x6a, func() base.Instruction {
		return fmul
	})

	factory.Factory.AddInstruction(0x6b, func() base.Instruction {
		return dmul
	})

}
