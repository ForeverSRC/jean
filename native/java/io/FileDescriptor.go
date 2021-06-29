package io

import (
	"jean/constants"
	"jean/native"
	"jean/rtda/jvmstack"
)

// private static native void initIDs();
func fdInitIDs(frame *jvmstack.Frame) {
	// todo
}

func init() {
	native.Register(constants.JavaIoFileDescriptor, "initIDs", "()V", fdInitIDs)
}
