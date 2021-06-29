package references

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (iv *INVOKE_SPECIAL) Execute(frame *jvmstack.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(iv.Index).(*heap.MethodRef)

	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()

	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}

	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	thisRef := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if thisRef == nil {
		panic("java.lang.NullPointerException")
	}

	if resolvedMethod.IsProtected() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		thisRef.Class() != currentClass &&
		!thisRef.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() && resolvedClass.IsSuperClassOf(currentClass) && resolvedMethod.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(), methodRef.Name(), methodRef.Descriptor())
	}

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}

func init() {
	factory.Factory.AddInstruction(0xb7, func() base.Instruction {
		return &INVOKE_SPECIAL{}
	})
}
