package references

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

type INSTANCE_OF struct {
	base.Index16Instruction
}

func (inst *INSTANCE_OF) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(inst.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}

func init() {
	factory.Factory.AddInstruction(0xc1, func() base.Instruction {
		return &INSTANCE_OF{}
	})
}
