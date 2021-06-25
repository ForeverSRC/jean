package misc

import (
	"jean/constants"
	"jean/instructions/base"
	"jean/native"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

//  private static native void initialize();
func initialize(frame *jvmstack.Frame) {
	// hack
	// private static native void initialize(){
	//	VM.savedProps.setProperty("foo","bar")
	//}
	vmClass := frame.Method().Class()
	savedProps := vmClass.GetRefVar("savedProps", "Ljava/util/Properties;")
	key := heap.JString(vmClass.Loader(), "foo")
	val := heap.JString(vmClass.Loader(), "bar")
	frame.OperandStack().PushRef(savedProps)

	frame.OperandStack().PushRef(key)
	frame.OperandStack().PushRef(val)

	propsClass := vmClass.Loader().LoadClass(constants.JavaUtilProperties)
	setPropMethod := propsClass.GetInstanceMethod("setProperty",
		"(Ljava/lang/String;Ljava/lang/String)Ljava/lang/Object;")
	base.InvokeMethod(frame, setPropMethod)
}

func init() {
	native.Registrer("sun/misc/VM", "initialize", "()V", initialize)
}
