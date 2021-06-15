package extended

import (
	"fmt"
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/instructions/loads"
	"jean/instructions/math"
	"jean/instructions/stores"
	"jean/rtda"
)

// WIDE 指令改变其他指令的行为，modifiedInstruction字段存放被改变的指令。
type WIDE struct {
	modifiedInstruction base.Instruction
}

var m = map[uint8]func(*base.BytecodeReader) base.Instruction{}

func init() {
	m[0x15] = func(reader *base.BytecodeReader) base.Instruction {
		ints := &loads.ILOAD{}
		ints.Index = uint(reader.ReadUint16())
		return ints
	}
	m[0x16] = func(reader *base.BytecodeReader) base.Instruction {
		ints := &loads.LLOAD{}
		ints.Index = uint(reader.ReadUint16())
		return ints
	}
	m[0x17] = func(reader *base.BytecodeReader) base.Instruction {
		ints := &loads.FLOAD{}
		ints.Index = uint(reader.ReadUint16())
		return ints
	}
	m[0x18] = func(reader *base.BytecodeReader) base.Instruction {
		ints := &loads.DLOAD{}
		ints.Index = uint(reader.ReadUint16())
		return ints
	}
	m[0x19] = func(reader *base.BytecodeReader) base.Instruction {
		ints := &loads.ALOAD{}
		ints.Index = uint(reader.ReadUint16())
		return ints
	}

	m[0x36] = func(reader *base.BytecodeReader) base.Instruction {
		ints := &stores.ISTORE{}
		ints.Index = uint(reader.ReadUint16())
		return ints
	}
	m[0x37] = func(reader *base.BytecodeReader) base.Instruction {
		ints := &stores.LSTORE{}
		ints.Index = uint(reader.ReadUint16())
		return ints
	}
	m[0x38] = func(reader *base.BytecodeReader) base.Instruction {
		ints := &stores.FSTORE{}
		ints.Index = uint(reader.ReadUint16())
		return ints
	}
	m[0x39] = func(reader *base.BytecodeReader) base.Instruction {
		ints := &stores.DSTORE{}
		ints.Index = uint(reader.ReadUint16())
		return ints
	}
	m[0x3a] = func(reader *base.BytecodeReader) base.Instruction {
		ints := &stores.ASTORE{}
		ints.Index = uint(reader.ReadUint16())
		return ints
	}

	m[0x84] = func(reader *base.BytecodeReader) base.Instruction {
		ints := &math.IINC{}
		ints.Index = uint(reader.ReadUint16())
		ints.Const = int32(reader.ReadInt16())
		return ints
	}

}

func (w *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	instFunc, ok := m[opcode]
	if !ok {
		panic(fmt.Errorf("unsupported opcode:%x", opcode))
	}

	w.modifiedInstruction = instFunc(reader)
}

func (w *WIDE) Execute(frame *rtda.Frame) {
	w.modifiedInstruction.Execute(frame)
}

func init() {
	factory.Factory.AddInstruction(0xc4, func() base.Instruction {
		return &WIDE{}
	})
}
