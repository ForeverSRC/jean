package lang

import (
	"jean/constants"
	"jean/native"
	"jean/rtda/jvmstack"
)

func init() {
	native.Register(constants.JavaLangThread, "currentThread", "()Ljava/lang/Thread;", currentThread)
	native.Register(constants.JavaLangThread, "setPriority0", "(I)V", setPriority0)
	native.Register(constants.JavaLangThread, "isAlive", "()Z", isAlive)
	native.Register(constants.JavaLangThread, "start0", "()V", start0)
}

// public static native Thread currentThread();
// ()Ljava/lang/Thread;
func currentThread(frame *jvmstack.Frame) {
	//jThread := frame.Thread().JThread()
	classLoader := frame.Method().Class().Loader()
	threadClass := classLoader.LoadClass("java/lang/Thread")
	jThread := threadClass.NewObject()

	threadGroupClass := classLoader.LoadClass("java/lang/ThreadGroup")
	jGroup := threadGroupClass.NewObject()

	jThread.SetRefVar("group", "Ljava/lang/ThreadGroup;", jGroup)
	jThread.SetIntVar("priority", "I", 1)

	frame.OperandStack().PushRef(jThread)
}

// private native void setPriority0(int newPriority);
// (I)V
func setPriority0(frame *jvmstack.Frame) {
	// vars := frame.LocalVars()
	// this := vars.GetThis()
	// newPriority := vars.GetInt(1))
	// todo
}

// public final native boolean isAlive();
// ()Z
func isAlive(frame *jvmstack.Frame) {
	//vars := frame.LocalVars()
	//this := vars.GetThis()

	stack := frame.OperandStack()
	stack.PushBoolean(false)
}

// private native void start0();
// ()V
func start0(frame *jvmstack.Frame) {
	// todo
}
