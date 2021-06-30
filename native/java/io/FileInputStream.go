package io

import (
	"jean/constants"
	"jean/native"
	"jean/rtda/jvmstack"
)

// private static native void initIDs();
func fisInitIDs(frame *jvmstack.Frame) {
	// todo
}

func init() {
	native.Register(constants.JavaIoFileInputStream, "initIDs", "()V", fisInitIDs)
}
