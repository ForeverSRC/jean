package references

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (a *ARRAY_LENGTH) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()

	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}

func init() {
	arraylength := &ARRAY_LENGTH{}
	factory.Factory.AddInstruction(0xbe, func() base.Instruction {
		return arraylength
	})
}
