package misc

import (
	"jean/constants"
	"jean/instructions/base"
	"jean/native"
	"jean/rtda/jvmstack"
)

//  private static native void initialize();
func initialize(frame *jvmstack.Frame) {

	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass(constants.JavaLangSystem)
	// private static void initializeSystemClass()
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}
