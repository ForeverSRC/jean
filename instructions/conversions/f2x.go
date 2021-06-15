package conversions

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type F2D struct {
	base.NoOperandsInstruction
}

func (f *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	target := float64(val)
	stack.PushDouble(target)
}

type F2I struct {
	base.NoOperandsInstruction
}

func (f *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	target := int32(val)
	stack.PushInt(target)
}

type F2L struct {
	base.NoOperandsInstruction
}

func (f *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	target := int64(val)
	stack.PushLong(target)
}

func init() {
	f2i := &F2I{}
	f2l := &F2L{}
	f2d := &F2D{}

	factory.Factory.AddInstruction(0x8b, func() base.Instruction {
		return f2i
	})

	factory.Factory.AddInstruction(0x8c, func() base.Instruction {
		return f2l
	})

	factory.Factory.AddInstruction(0x8d, func() base.Instruction {
		return f2d
	})

}
