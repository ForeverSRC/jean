package math

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type DNEG struct {
	base.NoOperandsInstruction
}

func (d *DNEG) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

type FNEG struct {
	base.NoOperandsInstruction
}

func (f *FNEG) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

type INEG struct {
	base.NoOperandsInstruction
}

func (i *INEG) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

type LNEG struct {
	base.NoOperandsInstruction
}

func (l *LNEG) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}

func init() {
	ineg := &INEG{}
	lneg := &LNEG{}
	fneg := &FNEG{}
	dneg := &DNEG{}

	factory.Factory.AddInstruction(0x74, func() base.Instruction {
		return ineg
	})

	factory.Factory.AddInstruction(0x75, func() base.Instruction {
		return lneg
	})

	factory.Factory.AddInstruction(0x76, func() base.Instruction {
		return fneg
	})

	factory.Factory.AddInstruction(0x77, func() base.Instruction {
		return dneg
	})

}
