package conversions

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type L2D struct {
	base.NoOperandsInstruction
}

func (l *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	target := float64(val)
	stack.PushDouble(target)
}

type L2F struct {
	base.NoOperandsInstruction
}

func (l *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	target := float32(val)
	stack.PushFloat(target)
}

type L2I struct {
	base.NoOperandsInstruction
}

func (l *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	target := int32(val)
	stack.PushInt(target)
}

func init() {
	l2i := &L2I{}
	l2f := &L2F{}
	l2d := &L2D{}

	factory.Factory.AddInstruction(0x88, func() base.Instruction {
		return l2i
	})

	factory.Factory.AddInstruction(0x89, func() base.Instruction {
		return l2f
	})

	factory.Factory.AddInstruction(0x8a, func() base.Instruction {
		return l2d
	})

}
