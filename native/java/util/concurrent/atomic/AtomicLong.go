package atomic

import (
	"jean/constants"
	"jean/native"
	"jean/rtda/jvmstack"
)

// private static native boolean VMSupportsCS8();
func vmSupportsCS8(frame *jvmstack.Frame) {
	frame.OperandStack().PushBoolean(false)
}

func init() {
	native.Register(constants.JUCAtomicAtomicLong, "VMSupportsCS8", "()Z", vmSupportsCS8)
}
