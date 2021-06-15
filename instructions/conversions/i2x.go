package conversions

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

// I2B Convert int to byte
type I2B struct {
	base.NoOperandsInstruction
}

func (i *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	b := int32(int8(val))
	stack.PushInt(b)
}

// I2C Convert int to char
type I2C struct {
	base.NoOperandsInstruction
}

func (i *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	c := int32(uint16(val))
	stack.PushInt(c)
}

type I2D struct {
	base.NoOperandsInstruction
}

func (i *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	target := float64(val)
	stack.PushDouble(target)
}

type I2F struct {
	base.NoOperandsInstruction
}

func (i *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	target := float32(val)
	stack.PushFloat(target)
}

type I2L struct {
	base.NoOperandsInstruction
}

func (i *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	target := int64(val)
	stack.PushLong(target)
}

// I2S Convert int to short
type I2S struct {
	base.NoOperandsInstruction
}

func (i *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	s := int32(int16(val))
	stack.PushInt(s)
}

func init() {
	i2l := &I2L{}
	i2f := &I2F{}
	i2d := &I2D{}

	i2b := &I2B{}
	i2c := &I2C{}
	i2s := &I2S{}

	factory.Factory.AddInstruction(0x85, func() base.Instruction {
		return i2l
	})

	factory.Factory.AddInstruction(0x86, func() base.Instruction {
		return i2f
	})

	factory.Factory.AddInstruction(0x87, func() base.Instruction {
		return i2d
	})

	factory.Factory.AddInstruction(0x91, func() base.Instruction {
		return i2b
	})

	factory.Factory.AddInstruction(0x92, func() base.Instruction {
		return i2c
	})

	factory.Factory.AddInstruction(0x93, func() base.Instruction {
		return i2s
	})
}
