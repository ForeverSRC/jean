package math

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
	"math"
)

type DREM struct {
	base.NoOperandsInstruction
}

func (d *DREM) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()

	res := math.Mod(v1, v2)
	stack.PushDouble(res)
}

type FREM struct {
	base.NoOperandsInstruction
}

func (f *FREM) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()

	res := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(res)
}

type IREM struct {
	base.NoOperandsInstruction
}

func (i *IREM) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	res := v1 % v2
	stack.PushInt(res)
}

type LREM struct {
	base.NoOperandsInstruction
}

func (l *LREM) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	res := v1 % v2
	stack.PushLong(res)
}

func init() {
	irem := &IREM{}
	lrem := &LREM{}
	frem := &FREM{}
	drem := &DREM{}

	factory.Factory.AddInstruction(0x70, func() base.Instruction {
		return irem
	})

	factory.Factory.AddInstruction(0x71, func() base.Instruction {
		return lrem
	})

	factory.Factory.AddInstruction(0x72, func() base.Instruction {
		return frem
	})

	factory.Factory.AddInstruction(0x73, func() base.Instruction {
		return drem
	})

}
