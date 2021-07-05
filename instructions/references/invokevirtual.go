package references

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (iv *INVOKE_VIRTUAL) Execute(frame *jvmstack.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(iv.Index).(*heap.MethodRef)

	// resolvedMethod 并不是最终要调用的方法，只是符号引用解析出的方法
	resolvedMethod := methodRef.ResolvedMethod()

	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	thisRef := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if thisRef == nil {
		// hack System.out.println
		panic("java.lang.NullPointerException")
	}

	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		thisRef.Class() != currentClass &&
		!thisRef.Class().IsSubClassOf(currentClass) {

		if !(thisRef.Class().IsArray() && resolvedMethod.Name() == "clone") {
			panic("java.lang.IllegalAccessError")
		}
	}

	var methodToBeInvoked *heap.Method
	if mtb := thisRef.Class().GetFromVtable(methodRef.Name(), methodRef.Descriptor()); mtb != nil {
		methodToBeInvoked = mtb
	} else {
		methodToBeInvoked = heap.LookupMethodInClass(thisRef.Class(), methodRef.Name(), methodRef.Descriptor())
		if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
			panic("java.lang.AbstractMethodError")
		}

		thisRef.Class().SetVtable(methodRef.Name(), methodRef.Descriptor(), methodToBeInvoked)
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}

func init() {
	factory.Factory.AddInstruction(0xb6, func() base.Instruction {
		return &INVOKE_VIRTUAL{}
	})
}
