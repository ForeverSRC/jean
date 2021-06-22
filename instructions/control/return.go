package control

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type RETURN struct {
	base.NoOperandsInstruction
}

func (r *RETURN) Execute(frame *jvmstack.Frame) {
	frame.Thread().PopFrame()
}

type ARETURN struct {
	base.NoOperandsInstruction
}

func (r *ARETURN) Execute(frame *jvmstack.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()

	retVal := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(retVal)
}

type DRETURN struct {
	base.NoOperandsInstruction
}

func (r *DRETURN) Execute(frame *jvmstack.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()

	retVal := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(retVal)
}

type FRETURN struct {
	base.NoOperandsInstruction
}

func (r *FRETURN) Execute(frame *jvmstack.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()

	retVal := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(retVal)
}

type IRETURN struct {
	base.NoOperandsInstruction
}

func (r *IRETURN) Execute(frame *jvmstack.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()

	retVal := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}

type LRETURN struct {
	base.NoOperandsInstruction
}

func (r *LRETURN) Execute(frame *jvmstack.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()

	retVal := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(retVal)
}

func init() {
	ireturn := &IRETURN{}
	lreturn := &LRETURN{}
	freturn := &FRETURN{}
	dreturn := &DRETURN{}
	areturn := &ARETURN{}
	_return := &RETURN{}

	factory.Factory.AddInstruction(0xac, func() base.Instruction {
		return ireturn
	})

	factory.Factory.AddInstruction(0xad, func() base.Instruction {
		return lreturn
	})

	factory.Factory.AddInstruction(0xae, func() base.Instruction {
		return freturn
	})

	factory.Factory.AddInstruction(0xaf, func() base.Instruction {
		return dreturn
	})

	factory.Factory.AddInstruction(0xb0, func() base.Instruction {
		return areturn
	})

	factory.Factory.AddInstruction(0xb1, func() base.Instruction {
		return _return
	})
}
