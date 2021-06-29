package lang

import (
	"jean/native"
	"jean/rtda/jvmstack"
	"runtime"
)

const jlRuntime = "java/lang/Runtime"

func init() {
	native.Register(jlRuntime, "availableProcessors", "()I", availableProcessors)
}

// public native int availableProcessors();
// ()I
func availableProcessors(frame *jvmstack.Frame) {
	numCPU := runtime.NumCPU()

	stack := frame.OperandStack()
	stack.PushInt(int32(numCPU))
}
