package reflect

import (
	"jean/constants"
	"jean/instructions/base"
	"jean/native"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

func init() {
	native.Register(constants.SunReflectNativeConstructorAccessorImpl, "newInstance0", "(Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;", newInstance0)
}

// private static native Object newInstance0(Constructor<?> c, Object[] os)
// throws InstantiationException, IllegalArgumentException, InvocationTargetException;
// (Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;
func newInstance0(frame *jvmstack.Frame) {
	vars := frame.LocalVars()
	constructorObj := vars.GetRef(0)
	argArrObj := vars.GetRef(1)

	goConstructor := getGoConstructor(constructorObj)
	goClass := goConstructor.Class()
	if !goClass.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), goClass)
		return
	}

	obj := goClass.NewObject()
	stack := frame.OperandStack()
	stack.PushRef(obj)

	// call <init>
	ops := convertArgs(obj, argArrObj, goConstructor)
	shimFrame := jvmstack.NewShimFrame(frame.Thread(), ops)
	frame.Thread().PushFrame(shimFrame)

	base.InvokeMethod(shimFrame, goConstructor)
}

func getGoMethod(methodObj *heap.Object) *heap.Method {
	return _getGoMethod(methodObj, false)
}
func getGoConstructor(constructorObj *heap.Object) *heap.Method {
	return _getGoMethod(constructorObj, true)
}
func _getGoMethod(methodObj *heap.Object, isConstructor bool) *heap.Method {
	extra := methodObj.Extra()
	if extra != nil {
		return extra.(*heap.Method)
	}

	if isConstructor {
		root := methodObj.GetRefVar("root", "Ljava/lang/reflect/Constructor;")
		return root.Extra().(*heap.Method)
	} else {
		root := methodObj.GetRefVar("root", "Ljava/lang/reflect/Method;")
		return root.Extra().(*heap.Method)
	}
}

// Object[] -> []interface{}
func convertArgs(this, argArr *heap.Object, method *heap.Method) *jvmstack.OperandStack {
	if method.ArgSlotCount() == 0 {
		return nil
	}

	//	argObjs := argArr.Refs()
	//	argTypes := method.ParsedDescriptor().ParameterTypes()

	ops := jvmstack.NewOperandStack(method.ArgSlotCount())
	if !method.IsStatic() {
		ops.PushRef(this)
	}
	if method.ArgSlotCount() == 1 && !method.IsStatic() {
		return ops
	}

	//	for i, argType := range argTypes {
	//		argObj := argObjs[i]
	//
	//		if len(argType) == 1 {
	//			// base type
	//			// todo
	//			unboxed := box.Unbox(argObj, argType)
	//			args[i+j] = unboxed
	//			if argType.isLongOrDouble() {
	//				j++
	//			}
	//		} else {
	//			args[i+j] = argObj
	//		}
	//	}

	return ops
}
