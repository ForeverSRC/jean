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

func init() {
	native.Registrer(constants.JavaLangDouble, "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
}
