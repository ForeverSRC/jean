package lang

import (
	"jean/constants"
	"jean/native"
	"jean/rtda/jvmstack"
	"math"
)

// public static native long doubleToRawLongBits(double value);
func doubleToRawLongBits(frame *jvmstack.Frame) {
	value := frame.LocalVars().GetDouble(0)
	bits := math.Float64bits(value)
	frame.OperandStack().PushLong(int64(bits))
}

// public static native double longBitsToDouble(long bits);
func longBitsToDouble(frame *jvmstack.Frame) {
	bits := frame.LocalVars().GetLong(0)
	value := math.Float64frombits(uint64(bits))
	frame.OperandStack().PushDouble(value)
}

func init() {
	native.Register(constants.JavaLangDouble, "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	native.Register(constants.JavaLangDouble, "longBitsToDouble", "(J)D", longBitsToDouble)
}
