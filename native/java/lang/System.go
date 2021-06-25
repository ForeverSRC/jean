package lang

import (
	"jean/constants"
	"jean/native"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

//public static native void arraycopy(Object src,  int  srcPos, Object dest, int destPos, int length);
// Javadoc:
//Throws:
//IndexOutOfBoundsException – if copying would cause access of data outside array bounds.
//ArrayStoreException – if an element in the src array could not be stored into the dest array because of a type mismatch.
//NullPointerException – if either src or dest is null.
func arrayCopy(frame *jvmstack.Frame) {
	vars := frame.LocalVars()

	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)

	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}

	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}

	if srcPos < 0 || destPos < 0 || length < 0 || srcPos+length > src.ArrayLength() || destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	heap.ArrayCopy(src, srcPos, dest, destPos, length)
}

func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()

	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}

	if srcClass.ComponentClass().IsPrimitive() || destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}

	return true
}

func init() {
	native.Registrer(constants.JavaLangSystem,
		"arraycopy",
		"(Ljava/long/Object;ILjava/lang/Object;II)V",
		arrayCopy)
}
