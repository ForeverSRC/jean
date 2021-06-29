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

// public static native float intBitsToFloat(int bits);
// (I)F
func intBitsToFloat(frame *jvmstack.Frame) {
	bits := frame.LocalVars().GetInt(0)
	value := math.Float32frombits(uint32(bits)) // todo
	frame.OperandStack().PushFloat(value)
}

func init() {
	native.Register(constants.JavaLangFloat, "floatToRawIntBits", "(F)I", floatToRawIntBits)
	native.Register(constants.JavaLangFloat, "intBitsToFloat", "(I)F", intBitsToFloat)
}
