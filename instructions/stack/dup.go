package stack

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/jvmstack"
)

/*
bottom -> top
[...][c][b][a]
             \_
               |
               V
[...][c][b][a][a]
*/
type DUP struct {
	base.NoOperandsInstruction
}

func (d *DUP) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

/*
bottom -> top
[...][c][b][a]
          __/
         |
         V
[...][c][a][b][a]
*/
type DUP_X1 struct {
	base.NoOperandsInstruction
}

func (d *DUP_X1) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()

	stack.PushSlot(slot1)

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
bottom -> top
[...][c][b][a]
       _____/
      |
      V
[...][a][c][b][a]
*/
type DUP_X2 struct {
	base.NoOperandsInstruction
}

func (d *DUP_X2) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()

	stack.PushSlot(slot1)

	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
bottom -> top
[...][c][b][a]____
          \____   |
               |  |
               V  V
[...][c][b][a][b][a]
*/
type DUP2 struct {
	base.NoOperandsInstruction
}

func (d *DUP2) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()

	for i := 0; i < 2; i++ {
		stack.PushSlot(slot2)
		stack.PushSlot(slot1)
	}
}

/*
bottom -> top
[...][c][b][a]
       _/ __/
      |  |
      V  V
[...][b][a][c][b][a]
*/
type DUP2_X1 struct {
	base.NoOperandsInstruction
}

func (d *DUP2_X1) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)

	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
bottom -> top
[...][d][c][b][a]
       ____/ __/
      |   __/
      V  V
[...][b][a][d][c][b][a]
*/
type DUP2_X2 struct {
	base.NoOperandsInstruction
}

func (d *DUP2_X2) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)

	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func init() {

	dup := &DUP{}
	dup_x1 := &DUP_X1{}
	dup_x2 := &DUP_X2{}

	dup2 := &DUP2{}
	dup2_x1 := &DUP2_X1{}
	dup2_x2 := &DUP2_X2{}

	factory.Factory.AddInstruction(0x59, func() base.Instruction {
		return dup
	})

	factory.Factory.AddInstruction(0x5a, func() base.Instruction {
		return dup_x1
	})

	factory.Factory.AddInstruction(0x5b, func() base.Instruction {
		return dup_x2
	})

	factory.Factory.AddInstruction(0x5c, func() base.Instruction {
		return dup2
	})

	factory.Factory.AddInstruction(0x5d, func() base.Instruction {
		return dup2_x1
	})

	factory.Factory.AddInstruction(0x5e, func() base.Instruction {
		return dup2_x2
	})

}
