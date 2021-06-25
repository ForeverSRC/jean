package lang

import (
	"jean/constants"
	"jean/native"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

// static native Class<?> getPrimitiveClass(String name);
func getPrimitiveClass(frame *jvmstack.Frame) {
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)

	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()

	frame.OperandStack().PushRef(class)
}

// private native String getName0();
func getName0(frame *jvmstack.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)

	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)

	frame.OperandStack().PushRef(nameObj)
}

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
func desiredAssertionStatus0(frame *jvmstack.Frame) {
	// todo 暂时不考虑断言
	frame.OperandStack().PushBoolean(false)
}

func init() {
	native.Registrer(constants.JavaLangClass,
		"getPrimitiveClass",
		"(Ljava/lang/String;)Ljava/lang/Class;",
		getPrimitiveClass)

	native.Registrer(constants.JavaLangClass,
		"getName0",
		"()Ljava/lang/String;",
		getName0)

	native.Registrer(constants.JavaLangClass,
		"desiredAssertionStatus0",
		"(Ljava/lang/Class;)Z",
		desiredAssertionStatus0)

}
