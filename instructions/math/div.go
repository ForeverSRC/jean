package math

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type DDIV struct {
	base.NoOperandsInstruction
}

func (d *DDIV) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopDouble()
	v2 := stack.PopDouble()

	res := v1 / v2

	stack.PushDouble(res)
}

type FDIV struct {
	base.NoOperandsInstruction
}

func (f *FDIV) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopFloat()
	v2 := stack.PopFloat()

	res := v1 / v2

	stack.PushFloat(res)
}

type IDIV struct {
	base.NoOperandsInstruction
}

func (i *IDIV) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopInt()
	v2 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	res := v1 / v2

	stack.PushInt(res)
}

type LDIV struct {
	base.NoOperandsInstruction
}

func (l *LDIV) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()

	v1 := stack.PopLong()
	v2 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	res := v1 / v2

	stack.PushLong(res)
}

func init() {
	idiv := &IDIV{}
	ldiv := &LDIV{}
	fdiv := &FDIV{}
	ddiv := &DDIV{}

	factory.Factory.AddInstruction(0x6c, func() base.Instruction {
		return idiv
	})

	factory.Factory.AddInstruction(0x6d, func() base.Instruction {
		return ldiv
	})

	factory.Factory.AddInstruction(0x6e, func() base.Instruction {
		return fdiv
	})

	factory.Factory.AddInstruction(0x6f, func() base.Instruction {
		return ddiv
	})

}
