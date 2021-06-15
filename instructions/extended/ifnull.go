package extended

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

type IFNULL struct {
	base.BranchInstruction
}

func (ifnull *IFNULL) Execute(frame *rtda.Frame) {
	_execute(frame, func(i interface{}) bool {
		return i == nil
	}, ifnull.Offset)
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (ifnonnull *IFNONNULL) Execute(frame *rtda.Frame) {
	_execute(frame, func(i interface{}) bool {
		return i != nil
	}, ifnonnull.Offset)
}

func _execute(frame *rtda.Frame, judgement func(interface{}) bool, offset int) {
	ref := frame.OperandStack().PopRef()
	if judgement(ref) {
		base.Branch(frame, offset)
	}
}

func init() {
	factory.Factory.AddInstruction(0xc6, func() base.Instruction {
		return &IFNULL{}
	})

	factory.Factory.AddInstruction(0xc7, func() base.Instruction {
		return &IFNONNULL{}
	})
}
