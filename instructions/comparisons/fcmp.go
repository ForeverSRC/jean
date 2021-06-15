package comparisons

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type FCMPG struct {
	base.NoOperandsInstruction
}

func (fcmpg *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (fcmpl *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()

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
	fcmpl := &FCMPL{}
	fcmpg := &FCMPG{}

	factory.Factory.AddInstruction(0x95, func() base.Instruction {
		return fcmpl
	})

	factory.Factory.AddInstruction(0x96, func() base.Instruction {
		return fcmpg
	})
}
