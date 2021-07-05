package references

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

type INVOKE_INTERFACE struct {
	index uint
	// count uint8
	// zero uint8
}

func (iv *INVOKE_INTERFACE) FetchOperands(reader *base.BytecodeReader) {
	iv.index = uint(reader.ReadUint16())
	//count
	reader.ReadUint8()
	// zero
	reader.ReadUint8()
}

func (iv *INVOKE_INTERFACE) Execute(frame *jvmstack.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(iv.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()

	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	thisRef := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	if thisRef == nil {
		panic("java.lang.NullPointerException")
	}

	if !thisRef.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	var methodToBeInvoked *heap.Method
	if mtb := thisRef.Class().GetFromItable(methodRef.Name(), methodRef.Descriptor()); mtb != nil {
		methodToBeInvoked = mtb
	} else {
		methodToBeInvoked = heap.LookupMethodInClass(thisRef.Class(), methodRef.Name(), methodRef.Descriptor())
		if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
			panic("java.lang.AbstractMethodError")
		}

		if !methodToBeInvoked.IsPublic() {
			panic("java.lang.IllegalAccessError")
		}

		thisRef.Class().SetItable(methodRef.Name(), methodRef.Descriptor(), methodToBeInvoked)
	}

	base.InvokeMethod(frame, methodToBeInvoked)
}

func init() {
	factory.Factory.AddInstruction(0xb9, func() base.Instruction {
		return &INVOKE_INTERFACE{}
	})
}
