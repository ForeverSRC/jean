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

// private static native long set(int d);
// (I)J
func set(frame *jvmstack.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}

func init() {
	native.Register(constants.JavaIoFileDescriptor, "initIDs", "()V", fdInitIDs)
	native.Register(constants.JavaIoFileDescriptor, "set", "(I)J", set)
}
