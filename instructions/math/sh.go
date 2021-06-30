package math

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type ISHL struct {
	base.NoOperandsInstruction
}

func (i *ISHL) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	// int变量只有32位，所以只取v2的前5个比特
	// Go语言位移操作符右侧必须是无符号整数，需要对v2进行类型转换
	s := uint32(v2) & 0x1f
	res := v1 << s

	stack.PushInt(res)
}

type ISHR struct {
	base.NoOperandsInstruction
}

func (i *ISHR) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	s := uint32(v2) & 0x1f
	res := v1 >> s

	stack.PushInt(res)
}

type IUSHR struct {
	base.NoOperandsInstruction
}

func (i *IUSHR) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	s := uint32(v2) & 0x1f
	res := int32(uint32(v1) >> s)

	stack.PushInt(res)
}

type LSHL struct {
	base.NoOperandsInstruction
}

func (i *LSHL) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()

	// Go语言位移操作符右侧必须是无符号整数，需要对v2进行类型转换
	s := uint32(v2) & 0x3f
	res := v1 << s

	stack.PushLong(res)
}

type LSHR struct {
	base.NoOperandsInstruction
}

func (l *LSHR) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()

	// long变量有64位，取v2前6个bit
	s := uint32(v2) & 0x3f
	res := v1 >> s

	stack.PushLong(res)
}

type LUSHR struct {
	base.NoOperandsInstruction
}

func (l *LUSHR) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()

	s := uint32(v2) & 0x3f
	res := int64(uint64(v1) >> s)

	stack.PushLong(res)
}

func init() {
	ishl := &ISHL{}
	lshl := &LSHL{}
	ishr := &ISHR{}
	lshr := &LSHR{}
	iushr := &IUSHR{}
	lushr := &LUSHR{}

	factory.Factory.AddInstruction(0x78, func() base.Instruction {
		return ishl
	})

	factory.Factory.AddInstruction(0x79, func() base.Instruction {
		return lshl
	})

	factory.Factory.AddInstruction(0x7a, func() base.Instruction {
		return ishr
	})

	factory.Factory.AddInstruction(0x7b, func() base.Instruction {
		return lshr
	})

	factory.Factory.AddInstruction(0x7c, func() base.Instruction {
		return iushr
	})

	factory.Factory.AddInstruction(0x7d, func() base.Instruction {
		return lushr
	})

}
