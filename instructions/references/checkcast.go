package references

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

type CHECK_CAST struct {
	base.Index16Instruction
}

func (cc *CHECK_CAST) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(cc.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}

func init() {
	factory.Factory.AddInstruction(0xc0, func() base.Instruction {
		return &CHECK_CAST{}
	})
}
