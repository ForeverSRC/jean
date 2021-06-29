package reflect

import (
	"jean/constants"
	"jean/native"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

// public static native Class<?> getCallerClass();
func getCallerClass(frame *jvmstack.Frame) {
	// top0 is sun/reflect/Reflection
	// top1 is the caller of getCallerClass()
	// top2 is the caller of method
	callerFrame := frame.Thread().GetFrames()[2] // todo
	callerClass := callerFrame.Method().Class().JClass()
	frame.OperandStack().PushRef(callerClass)
}

// public static native int getClassAccessFlags(Class<?> type);
func getClassAccessFlags(frame *jvmstack.Frame) {
	vars := frame.LocalVars()
	_type := vars.GetRef(0)

	goClass := _type.Extra().(*heap.Class)
	flags := goClass.AccessFlags()

	stack := frame.OperandStack()
	stack.PushInt(int32(flags))
}

func init() {
	native.Register(constants.SunReflectReflection, "getCallerClass", "()Ljava/lang/Class;", getCallerClass)
	native.Register(constants.SunReflectReflection, "getClassAccessFlags", "(Ljava/lang/Class;)I", getClassAccessFlags)
}
