package comparisons

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (ifIcmp *IF_ICMPEQ) Execute(frame *jvmstack.Frame) {
	_ifIcmp(frame, func(v1, v2 int32) bool {
		return v1 == v2
	}, ifIcmp.Offset)
}

type IF_ICMPNE struct {
	base.BranchInstruction
}

func (ifIcmp *IF_ICMPNE) Execute(frame *jvmstack.Frame) {
	_ifIcmp(frame, func(v1, v2 int32) bool {
		return v1 != v2
	}, ifIcmp.Offset)
}

type IF_ICMPLT struct {
	base.BranchInstruction
}

func (ifIcmp *IF_ICMPLT) Execute(frame *jvmstack.Frame) {
	_ifIcmp(frame, func(v1, v2 int32) bool {
		return v1 < v2
	}, ifIcmp.Offset)
}

type IF_ICMPLE struct {
	base.BranchInstruction
}

func (ifIcmp *IF_ICMPLE) Execute(frame *jvmstack.Frame) {
	_ifIcmp(frame, func(v1, v2 int32) bool {
		return v1 <= v2
	}, ifIcmp.Offset)
}

type IF_ICMPGT struct {
	base.BranchInstruction
}

func (ifIcmp *IF_ICMPGT) Execute(frame *jvmstack.Frame) {
	_ifIcmp(frame, func(v1, v2 int32) bool {
		return v1 > v2
	}, ifIcmp.Offset)
}

type IF_ICMPGE struct {
	base.BranchInstruction
}

func (ifIcmp *IF_ICMPGE) Execute(frame *jvmstack.Frame) {
	_ifIcmp(frame, func(v1, v2 int32) bool {
		return v1 >= v2
	}, ifIcmp.Offset)
}

func _ifIcmp(frame *jvmstack.Frame, cond func(int32, int32) bool, offset int) {
	val2 := frame.OperandStack().PopInt()
	val1 := frame.OperandStack().PopInt()
	if cond(val1, val2) {
		base.Branch(frame, offset)
	}
}

func init() {
	factory.Factory.AddInstruction(0x9f, func() base.Instruction {
		return &IF_ICMPEQ{}
	})

	factory.Factory.AddInstruction(0xa0, func() base.Instruction {
		return &IF_ICMPNE{}
	})

	factory.Factory.AddInstruction(0xa1, func() base.Instruction {
		return &IF_ICMPLT{}
	})

	factory.Factory.AddInstruction(0xa2, func() base.Instruction {
		return &IF_ICMPGE{}
	})

	factory.Factory.AddInstruction(0xa3, func() base.Instruction {
		return &IF_ICMPGT{}
	})

	factory.Factory.AddInstruction(0xa4, func() base.Instruction {
		return &IF_ICMPLE{}
	})
}
