package math

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type DSUB struct {
	base.NoOperandsInstruction
}

func (d *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopDouble()
	v2 := stack.PopDouble()

	res := v1 - v2

	stack.PushDouble(res)
}

type FSUB struct {
	base.NoOperandsInstruction
}

func (f *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopFloat()
	v2 := stack.PopFloat()

	res := v1 - v2

	stack.PushFloat(res)
}

type ISUB struct {
	base.NoOperandsInstruction
}

func (i *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopInt()
	v2 := stack.PopInt()

	res := v1 - v2

	stack.PushInt(res)
}

type LSUB struct {
	base.NoOperandsInstruction
}

func (l *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopLong()
	v2 := stack.PopLong()

	res := v1 - v2

	stack.PushLong(res)
}

func init() {
	isub := &ISUB{}
	lsub := &LSUB{}
	fsub := &FSUB{}
	dsub := &DSUB{}

	factory.Factory.AddInstruction(0x64, func() base.Instruction {
		return isub
	})

	factory.Factory.AddInstruction(0x65, func() base.Instruction {
		return lsub
	})

	factory.Factory.AddInstruction(0x66, func() base.Instruction {
		return fsub
	})

	factory.Factory.AddInstruction(0x67, func() base.Instruction {
		return dsub
	})

}
