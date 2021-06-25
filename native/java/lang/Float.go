package lang

import (
	"jean/constants"
	"jean/native"
	"jean/rtda/jvmstack"
	"math"
)

// public static native int floatToRawIntBits(float value);
func floatToRawIntBits(frame *jvmstack.Frame) {
	value := frame.LocalVars().GetFloat(0)
	bits := math.Float32bits(value)
	frame.OperandStack().PushInt(int32(bits))
}

func init() {
	native.Registrer(constants.JavaLangFloat, "floatToRawIntBits", "(F)I", floatToRawIntBits)
}
