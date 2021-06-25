package lang

import (
	"jean/constants"
	"jean/native"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

// javadoc:
// When the intern method is invoked,
// if the pool already contains a string equal to this String object as determined by the equals(Object) method,
// then the string from the pool is returned.
// Otherwise, this String object is added to the pool and a reference to this String object is returned.

// public native String intern();
func intern(frame *jvmstack.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}

func init() {
	native.Registrer(constants.JavaLangString, "intern", "()Ljava/lang/String;", intern)
}
