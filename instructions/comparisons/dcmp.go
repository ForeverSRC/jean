package comparisons

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type DCMPG struct {
	base.NoOperandsInstruction
}

func (dcmpg *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func (dcmpl *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()

	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func init() {
	dcmpl := &DCMPL{}
	dcmpg := &DCMPG{}

	factory.Factory.AddInstruction(0x97, func() base.Instruction {
		return dcmpl
	})

	factory.Factory.AddInstruction(0x98, func() base.Instruction {
		return dcmpg
	})
}
