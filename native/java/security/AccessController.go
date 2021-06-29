package security

import (
	"jean/constants"
	"jean/instructions/base"
	"jean/native"
	"jean/rtda/jvmstack"
)

// @CallerSensitive
// public static native <T> T doPrivileged(PrivilegedExceptionAction<T> action)  throws PrivilegedActionException;
// (Ljava/security/PrivilegedExceptionAction;)Ljava/lang/Object;

// @CallerSensitive
// public static native <T> T doPrivileged(PrivilegedAction<T> action);
// (Ljava/security/PrivilegedAction;)Ljava/lang/Object;
func doPrivileged(frame *jvmstack.Frame) {
	vars := frame.LocalVars()
	action := vars.GetRef(0)

	stack := frame.OperandStack()
	stack.PushRef(action)

	method := action.Class().GetInstanceMethod("run", "()Ljava/lang/Object;") // todo
	base.InvokeMethod(frame, method)
}

// private static native AccessControlContext getStackAccessControlContext();
// ()Ljava/security/AccessControlContext;
func getStackAccessControlContext(frame *jvmstack.Frame) {
	// todo
	frame.OperandStack().PushRef(nil)
}

func init() {
	native.Register(constants.JavaSecurityAccessController,
		"doPrivileged",
		"(Ljava/security/PrivilegedExceptionAction;)Ljava/lang/Object;",
		doPrivileged)
	native.Register(constants.JavaSecurityAccessController,
		"doPrivileged",
		"(Ljava/security/PrivilegedAction;)Ljava/lang/Object;",
		doPrivileged)
	native.Register(constants.JavaSecurityAccessController,
		"getStackAccessControlContext",
		"()Ljava/security/AccessControlContext;",
		getStackAccessControlContext)
}
