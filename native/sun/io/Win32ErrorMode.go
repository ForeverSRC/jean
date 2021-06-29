package io

import (
	"jean/native"
	"jean/rtda/jvmstack"
)

func init() {
	native.Register("sun/io/Win32ErrorMode", "setErrorMode", "(J)J", setErrorMode)
}

func setErrorMode(frame *jvmstack.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}
