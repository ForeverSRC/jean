package loads

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

type AALOAD struct {
	base.NoOperandsInstruction
}

func (al *AALOAD) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	refs := arrRef.Refs()

	checkIndex(len(refs), index)
	stack.PushRef(refs[index])
}

type BALOAD struct {
	base.NoOperandsInstruction
}

func (bl *BALOAD) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	bytes := arrRef.Bytes()

	checkIndex(len(bytes), index)
	stack.PushInt(int32(bytes[index]))
}

type CALOAD struct {
	base.NoOperandsInstruction
}

func (cl *CALOAD) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	chars := arrRef.Chars()

	checkIndex(len(chars), index)
	stack.PushInt(int32(chars[index]))
}

type DALOAD struct {
	base.NoOperandsInstruction
}

func (dl *DALOAD) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	doubles := arrRef.Doubles()

	checkIndex(len(doubles), index)
	stack.PushDouble(doubles[index])
}

type FALOAD struct {
	base.NoOperandsInstruction
}

func (fl *FALOAD) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	floats := arrRef.Floats()

	checkIndex(len(floats), index)
	stack.PushFloat(floats[index])
}

type IALOAD struct {
	base.NoOperandsInstruction
}

func (il *IALOAD) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	ints := arrRef.Ints()

	checkIndex(len(ints), index)
	stack.PushInt(ints[index])
}

type LALOAD struct {
	base.NoOperandsInstruction
}

func (ll *LALOAD) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	longs := arrRef.Longs()

	checkIndex(len(longs), index)
	stack.PushLong(longs[index])
}

type SALOAD struct {
	base.NoOperandsInstruction
}

func (sl *SALOAD) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	shorts := arrRef.Shorts()

	checkIndex(len(shorts), index)
	stack.PushInt(int32(shorts[index]))
}

func checkNotNull(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("java.lang.ArrayIndexOutOfBoundException")
	}
}

func init() {
	aaload := &AALOAD{}
	baload := &BALOAD{}
	caload := &CALOAD{}
	daload := &DALOAD{}
	faload := &FALOAD{}
	iaload := &IALOAD{}
	laload := &LALOAD{}
	saload := &SALOAD{}

	factory.Factory.AddInstruction(0x2e, func() base.Instruction {
		return iaload
	})

	factory.Factory.AddInstruction(0x2f, func() base.Instruction {
		return laload
	})

	factory.Factory.AddInstruction(0x30, func() base.Instruction {
		return faload
	})

	factory.Factory.AddInstruction(0x31, func() base.Instruction {
		return daload
	})

	factory.Factory.AddInstruction(0x32, func() base.Instruction {
		return aaload
	})

	factory.Factory.AddInstruction(0x33, func() base.Instruction {
		return baload
	})

	factory.Factory.AddInstruction(0x34, func() base.Instruction {
		return caload
	})

	factory.Factory.AddInstruction(0x35, func() base.Instruction {
		return saload
	})
}
