package constants

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
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

	switch ct := c.(type) {
	case int64:
		stack.PushLong(ct)
	case float64:
		stack.PushDouble(ct)
	default:
		panic("java.lang.ClassFormatError")
	}
}

func _ldc(frame *jvmstack.Frame, index uint) {
	stack := frame.OperandStack()
	class := frame.Method().Class()
	c := class.ConstantPool().GetConstant(index)

	switch ct := c.(type) {
	case int32:
		stack.PushInt(ct)
	case float32:
		stack.PushFloat(ct)
	case string:
		internedStr := heap.JString(class.Loader(), ct)
		stack.PushRef(internedStr)
	case *heap.ClassRef:
		classObj := ct.ResolvedClass().JClass()
		stack.PushRef(classObj)
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
