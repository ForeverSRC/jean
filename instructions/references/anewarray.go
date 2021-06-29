package references

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

type ANEW_ARRAY struct {
	base.Index16Instruction
}

func (na *ANEW_ARRAY) Execute(frame *jvmstack.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(na.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	count := stack.PopInt()

	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

func init() {
	factory.Factory.AddInstruction(0xbd, func() base.Instruction {
		return &ANEW_ARRAY{}
	})
}
