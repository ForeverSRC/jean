package math

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type DMUL struct {
	base.NoOperandsInstruction
}

func (d *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopDouble()
	v2 := stack.PopDouble()

	res := v1 * v2

	stack.PushDouble(res)
}

type FMUL struct {
	base.NoOperandsInstruction
}

func (f *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopFloat()
	v2 := stack.PopFloat()

	res := v1 * v2

	stack.PushFloat(res)
}

type IMUL struct {
	base.NoOperandsInstruction
}

func (i *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopInt()
	v2 := stack.PopInt()

	res := v1 * v2

	stack.PushInt(res)
}

type LMUL struct {
	base.NoOperandsInstruction
}

func (l *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopLong()
	v2 := stack.PopLong()

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
