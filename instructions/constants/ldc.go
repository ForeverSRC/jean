package constants

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type LDC struct {
	base.Index8Instruction
}

func (ldc *LDC) Execute(frame *jvmstack.Frame) {
	_ldc(frame, ldc.Index)
}

type LDC_W struct {
	base.Index16Instruction
}

func (ldc *LDC_W) Execute(frame *jvmstack.Frame) {
	_ldc(frame, ldc.Index)
}

type LDC2_W struct {
	base.Index16Instruction
}

func (ldc *LDC2_W) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(ldc.Index)

	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}

func _ldc(frame *jvmstack.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	// case string
	// case *heap.ClassRef
	default:
		panic("todo: ldc!")
	}
}

func init() {
	factory.Factory.AddInstruction(0x12, func() base.Instruction {
		return &LDC{}
	})

	factory.Factory.AddInstruction(0x13, func() base.Instruction {
		return &LDC_W{}
	})

	factory.Factory.AddInstruction(0x14, func() base.Instruction {
		return &LDC2_W{}
	})
}
