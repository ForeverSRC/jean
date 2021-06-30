package native

import (
	"fmt"
	"jean/rtda/jvmstack"
)

type NativeMethod func(frame *jvmstack.Frame)

var registry = map[string]NativeMethod{}

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := generateKey(className, methodName, methodDescriptor)
	registry[key] = method
}

func FindNativeMethod(classname, methodName, methodDescriptor string) NativeMethod {
	key := generateKey(classname, methodName, methodDescriptor)

	if method, ok := registry[key]; ok {
		return method
	}

	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}

	return nil
}

func generateKey(className, methodName, methodDescriptor string) string {
	return fmt.Sprintf("%s~%s~%s", className, methodName, methodDescriptor)
}

func emptyNativeMethod(frame *jvmstack.Frame) {
	// do nothing
}
