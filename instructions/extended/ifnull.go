package extended

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type IFNULL struct {
	base.BranchInstruction
}

func (ifnull *IFNULL) Execute(frame *jvmstack.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, ifnull.Offset)
	}
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (ifnonnull *IFNONNULL) Execute(frame *jvmstack.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, ifnonnull.Offset)
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
