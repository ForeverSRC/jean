package references

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

type NEW struct {
	base.Index16Instruction
}

func (n *NEW) Execute(frame *jvmstack.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(n.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}

func init() {
	factory.Factory.AddInstruction(0xbb, func() base.Instruction {
		return &NEW{}
	})
}
