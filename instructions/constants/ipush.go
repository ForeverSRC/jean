package constants

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

// BIPUSH 从操作数中获取一个byte整数，扩展成int并推入栈顶
type BIPUSH struct {
	val int8
}

func (bipush *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	bipush.val = reader.ReadInt8()
}

func (bipush *BIPUSH) Execute(frame *jvmstack.Frame) {
	i := int32(bipush.val)
	frame.OperandStack().PushInt(i)
}

// SIPUSH 从操作数中获取一个short整数，扩展成int并推入栈顶
type SIPUSH struct {
	val int16
}

func (sipush *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	sipush.val = reader.ReadInt16()
}

func (sipush *SIPUSH) Execute(frame *jvmstack.Frame) {
	i := int32(sipush.val)
	frame.OperandStack().PushInt(i)
}

func init() {
	factory.Factory.AddInstruction(0x10, func() base.Instruction {
		return &BIPUSH{}
	})

	factory.Factory.AddInstruction(0x11, func() base.Instruction {
		return &SIPUSH{}
	})
}
