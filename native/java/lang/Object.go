package lang

import (
	"jean/constants"
	"jean/native"
	"jean/rtda/jvmstack"
	"unsafe"
)

// public final native Class<?> getClass();
func getClass(frame *jvmstack.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}

// public native int hashCode();
func hashCode(frame *jvmstack.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

// protected native Object clone() throws CloneNotSupportedException;
func clone(frame *jvmstack.Frame) {
	this := frame.LocalVars().GetThis()
	cloneable := this.Class().Loader().LoadClass(constants.JavaLangCloneable)
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}

	frame.OperandStack().PushRef(this.Clone())
}

func init() {
	native.Registrer(constants.JavaLangObject, "getClass", "()Ljava/lang/Class;", getClass)
	native.Registrer(constants.JavaLangObject, "hashCode", "()I", hashCode)
	native.Registrer(constants.JavaLangObject, "clone", "()Ljava/lang/Object;", clone)
}
