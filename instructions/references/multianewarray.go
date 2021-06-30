package references

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
}

func (ma *MULTI_ANEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	ma.index = reader.ReadUint16()
	ma.dimensions = reader.ReadUint8()
}

func (ma *MULTI_ANEW_ARRAY) Execute(frame *jvmstack.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(ma.index)).(*heap.ClassRef)
	arrClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	counts := popAndCheckCounts(stack, int(ma.dimensions))
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
}

func popAndCheckCounts(stack *jvmstack.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for i := dimensions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}

	return counts
}

func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := arrClass.NewArray(count)

	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}

	return arr
}

func init() {
	factory.Factory.AddInstruction(0xc5, func() base.Instruction {
		return &MULTI_ANEW_ARRAY{}
	})
}
