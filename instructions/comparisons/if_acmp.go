package comparisons

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (ifAcmp *IF_ACMPEQ) Execute(frame *jvmstack.Frame) {
	_ifAcmp(frame, func(r1, r2 *heap.Object) bool {
		return r1 == r2
	}, ifAcmp.Offset)
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (ifAcmp *IF_ACMPNE) Execute(frame *jvmstack.Frame) {
	_ifAcmp(frame, func(r1, r2 *heap.Object) bool {
		return r1 != r2
	}, ifAcmp.Offset)
}

func _ifAcmp(frame *jvmstack.Frame, cond func(r1, r2 *heap.Object) bool, offset int) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()

	if cond(ref1, ref2) {
		base.Branch(frame, offset)
	}
}

func init() {
	factory.Factory.AddInstruction(0xa5, func() base.Instruction {
		return &IF_ACMPEQ{}
	})

	factory.Factory.AddInstruction(0xa6, func() base.Instruction {
		return &IF_ACMPNE{}
	})
}
