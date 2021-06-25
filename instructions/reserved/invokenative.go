package reserved

import (
	"fmt"
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/native"
	_ "jean/native/java/lang"
	_ "jean/native/sun/misc"
	"jean/rtda/jvmstack"
)

type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (iv *INVOKE_NATIVE) Execute(frame *jvmstack.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		panic("java.lang.UnsatisfiedLinkError: " + fmt.Sprintf("%s.%s%s", className, methodName, methodDescriptor))
	}

	nativeMethod(frame)
}

func init() {
	invokeNative := &INVOKE_NATIVE{}
	factory.Factory.AddInstruction(0xfe, func() base.Instruction {
		return invokeNative
	})
}
