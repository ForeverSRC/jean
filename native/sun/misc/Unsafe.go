package misc

import (
	"jean/constants"
	"jean/native"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

// public native int arrayBaseOffset(Class<?> type);
func arrayBaseOffset(frame *jvmstack.Frame) {
	frame.OperandStack().PushInt(0)
}

// public native int arrayIndexScale(Class<?> type);
func arrayIndexScale(frame *jvmstack.Frame) {
	frame.OperandStack().PushInt(1)
}

// public native int addressSize();
func addressSize(frame *jvmstack.Frame) {
	frame.OperandStack().PushInt(8)
}

// public native long objectFieldOffset(Field field);
func objectFieldOffset(frame *jvmstack.Frame) {
	vars := frame.LocalVars()
	jField := vars.GetRef(1)

	offset := jField.GetIntVar("slot", "I")
	stack := frame.OperandStack()
	stack.PushLong(int64(offset))
}

// public final native boolean compareAndSwapObject(Object o, long offset, Object expected, Object x)
// (Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z
func compareAndSwapObject(frame *jvmstack.Frame) {
	vars := frame.LocalVars()
	obj := vars.GetRef(1)
	fields := obj.Data()
	offset := vars.GetLong(2)
	expected := vars.GetRef(4)
	newVal := vars.GetRef(5)

	// todo
	if anys, ok := fields.(heap.Slots); ok {
		// object
		swapped := _casObj(obj, anys, offset, expected, newVal)
		frame.OperandStack().PushBoolean(swapped)
	} else if objs, ok := fields.([]*heap.Object); ok {
		// ref[]
		swapped := _casArr(objs, offset, expected, newVal)
		frame.OperandStack().PushBoolean(swapped)
	} else {
		// todo
		panic("todo: compareAndSwapObject!")
	}
}
func _casObj(obj *heap.Object, fields heap.Slots, offset int64, expected, newVal *heap.Object) bool {
	current := fields.GetRef(uint(offset))
	if current == expected {
		fields.SetRef(uint(offset), newVal)
		return true
	} else {
		return false
	}
}
func _casArr(objs []*heap.Object, offset int64, expected, newVal *heap.Object) bool {
	current := objs[offset]
	if current == expected {
		objs[offset] = newVal
		return true
	} else {
		return false
	}
}

// public native boolean getInt(Object o, long offset);
// (Ljava/lang/Object;J)I
func getInt(frame *jvmstack.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Data()
	offset := vars.GetLong(2)

	stack := frame.OperandStack()
	if slots, ok := fields.(heap.Slots); ok {
		// object
		stack.PushInt(slots.GetInt(uint(offset)))
	} else if shorts, ok := fields.([]int32); ok {
		// int[]
		stack.PushInt(int32(shorts[offset]))
	} else {
		panic("getInt!")
	}
}

// public final native boolean compareAndSwapInt(Object o, long offset, int expected, int x);
// (Ljava/lang/Object;JII)Z
func compareAndSwapInt(frame *jvmstack.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Data()
	offset := vars.GetLong(2)
	expected := vars.GetInt(4)
	newVal := vars.GetInt(5)

	if slots, ok := fields.(heap.Slots); ok {
		// object
		oldVal := slots.GetInt(uint(offset))
		if oldVal == expected {
			slots.SetInt(uint(offset), newVal)
			frame.OperandStack().PushBoolean(true)
		} else {
			frame.OperandStack().PushBoolean(false)
		}
	} else if ints, ok := fields.([]int32); ok {
		// int[]
		oldVal := ints[offset]
		if oldVal == expected {
			ints[offset] = newVal
			frame.OperandStack().PushBoolean(true)
		} else {
			frame.OperandStack().PushBoolean(false)
		}
	} else {
		// todo
		panic("todo: compareAndSwapInt!")
	}
}

// public native Object getObject(Object o, long offset);
// (Ljava/lang/Object;J)Ljava/lang/Object;
func getObject(frame *jvmstack.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Data()
	offset := vars.GetLong(2)

	if anys, ok := fields.(heap.Slots); ok {
		// object
		x := anys.GetRef(uint(offset))
		frame.OperandStack().PushRef(x)
	} else if objs, ok := fields.([]*heap.Object); ok {
		// ref[]
		x := objs[offset]
		frame.OperandStack().PushRef(x)
	} else {
		panic("getObject!")
	}
}

// public final native boolean compareAndSwapLong(Object o, long offset, long expected, long x);
// (Ljava/lang/Object;JJJ)Z
func compareAndSwapLong(frame *jvmstack.Frame) {
	vars := frame.LocalVars()
	fields := vars.GetRef(1).Data()
	offset := vars.GetLong(2)
	expected := vars.GetLong(4)
	newVal := vars.GetLong(6)

	if slots, ok := fields.(heap.Slots); ok {
		// object
		oldVal := slots.GetLong(uint(offset))
		if oldVal == expected {
			slots.SetLong(uint(offset), newVal)
			frame.OperandStack().PushBoolean(true)
		} else {
			frame.OperandStack().PushBoolean(false)
		}
	} else if longs, ok := fields.([]int64); ok {
		// long[]
		oldVal := longs[offset]
		if oldVal == expected {
			longs[offset] = newVal
			frame.OperandStack().PushBoolean(true)
		} else {
			frame.OperandStack().PushBoolean(false)
		}
	} else {
		// todo
		panic("todo: compareAndSwapLong!")
	}
}

func init() {
	native.Register(constants.SunMiscUnsafe, "arrayBaseOffset", "(Ljava/lang/Class;)I", arrayBaseOffset)
	native.Register(constants.SunMiscUnsafe, "arrayIndexScale", "(Ljava/lang/Class;)I", arrayIndexScale)
	native.Register(constants.SunMiscUnsafe, "addressSize", "()I", addressSize)
	native.Register(constants.SunMiscUnsafe, "objectFieldOffset", "(Ljava/lang/reflect/Field;)J", objectFieldOffset)
	native.Register(constants.SunMiscUnsafe, "compareAndSwapObject", "(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z", compareAndSwapObject)
	native.Register(constants.SunMiscUnsafe, "getIntVolatile", "(Ljava/lang/Object;J)I", getInt)
	native.Register(constants.SunMiscUnsafe, "compareAndSwapInt", "(Ljava/lang/Object;JII)Z", compareAndSwapInt)
	native.Register(constants.SunMiscUnsafe, "getObjectVolatile", "(Ljava/lang/Object;J)Ljava/lang/Object;", getObject)
	native.Register(constants.SunMiscUnsafe, "compareAndSwapLong", "(Ljava/lang/Object;JJJ)Z", compareAndSwapLong)

}
