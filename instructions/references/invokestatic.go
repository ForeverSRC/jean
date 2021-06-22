package references

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

// invoke a class static method
type INVOKE_STATIC struct {
	base.Index16Instruction
}

func (iv *INVOKE_STATIC) Execute(frame *jvmstack.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodref := cp.GetConstant(iv.Index).(*heap.MethodRef)
	resolvedMethod := methodref.ResolvedMethod()

	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	class := resolvedMethod.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	base.InvokeMethod(frame, resolvedMethod)
}

func init() {
	factory.Factory.AddInstruction(0xb8, func() base.Instruction {
		return &INVOKE_STATIC{}
	})
}
