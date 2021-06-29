package stores

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

type AASTORE struct {
	base.NoOperandsInstruction
}

func (as *AASTORE) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := stack.PopRef()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	refs := arrRef.Refs()

	checkIndex(len(refs), index)
	refs[index] = val
}

type BASTORE struct {
	base.NoOperandsInstruction
}

func (bs *BASTORE) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := int8(stack.PopInt())
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	bytes := arrRef.Bytes()

	checkIndex(len(bytes), index)
	bytes[index] = val
}

type CASTORE struct {
	base.NoOperandsInstruction
}

func (cs *CASTORE) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := uint16(stack.PopInt())
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	chars := arrRef.Chars()

	checkIndex(len(chars), index)
	chars[index] = val
}

type DASTORE struct {
	base.NoOperandsInstruction
}

func (ds *DASTORE) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	doubles := arrRef.Doubles()

	checkIndex(len(doubles), index)
	doubles[index] = val
}

type FASTORE struct {
	base.NoOperandsInstruction
}

func (fs *FASTORE) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	floats := arrRef.Floats()

	checkIndex(len(floats), index)
	floats[index] = val
}

type IASTORE struct {
	base.NoOperandsInstruction
}

func (is *IASTORE) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	ints := arrRef.Ints()

	checkIndex(len(ints), index)
	ints[index] = val
}

type LASTORE struct {
	base.NoOperandsInstruction
}

func (ls *LASTORE) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	longs := arrRef.Longs()

	checkIndex(len(longs), index)
	longs[index] = val
}

type SASTORE struct {
	base.NoOperandsInstruction
}

func (ss *SASTORE) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	val := int16(stack.PopInt())
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNull(arrRef)
	shorts := arrRef.Shorts()

	checkIndex(len(shorts), index)
	shorts[index] = val
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
	aastore := &AASTORE{}
	bastore := &BASTORE{}
	castore := &CASTORE{}
	dastore := &DASTORE{}
	fastore := &FASTORE{}
	iastore := &IASTORE{}
	lastore := &LASTORE{}
	sastore := &SASTORE{}

	factory.Factory.AddInstruction(0x4f, func() base.Instruction {
		return iastore
	})

	factory.Factory.AddInstruction(0x50, func() base.Instruction {
		return lastore
	})

	factory.Factory.AddInstruction(0x51, func() base.Instruction {
		return fastore
	})

	factory.Factory.AddInstruction(0x52, func() base.Instruction {
		return dastore
	})

	factory.Factory.AddInstruction(0x53, func() base.Instruction {
		return aastore
	})

	factory.Factory.AddInstruction(0x54, func() base.Instruction {
		return bastore
	})

	factory.Factory.AddInstruction(0x55, func() base.Instruction {
		return castore
	})

	factory.Factory.AddInstruction(0x56, func() base.Instruction {
		return sastore
	})
}
