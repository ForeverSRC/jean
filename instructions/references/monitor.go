package references

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

type MONITOR_ENTER struct {
	base.NoOperandsInstruction
}

// todo
func (m *MONITOR_ENTER) Execute(frame *jvmstack.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

type MONITOR_EXIT struct {
	base.NoOperandsInstruction
}

// todo
func (m *MONITOR_EXIT) Execute(frame *jvmstack.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func init() {
	monitorenter := &MONITOR_ENTER{}
	monitorexit := &MONITOR_EXIT{}

	factory.Factory.AddInstruction(0xc2, func() base.Instruction {
		return monitorenter
	})

	factory.Factory.AddInstruction(0xc3, func() base.Instruction {
		return monitorexit
	})
}
