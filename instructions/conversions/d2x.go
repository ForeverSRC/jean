package conversions

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type D2F struct {
	base.NoOperandsInstruction
}

func (d *D2F) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	target := float32(val)
	stack.PushFloat(target)
}

type D2I struct {
	base.NoOperandsInstruction
}

func (d *D2I) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	target := int32(val)
	stack.PushInt(target)
}

type D2L struct {
	base.NoOperandsInstruction
}

func (d *D2L) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	target := int64(val)
	stack.PushLong(target)
}

func init() {
	d2i := &D2I{}
	d2l := &D2L{}
	d2f := &D2F{}

	factory.Factory.AddInstruction(0x8e, func() base.Instruction {
		return d2i
	})

	factory.Factory.AddInstruction(0x8f, func() base.Instruction {
		return d2l
	})

	factory.Factory.AddInstruction(0x90, func() base.Instruction {
		return d2f
	})

}
