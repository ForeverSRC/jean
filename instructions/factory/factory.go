package factory

import (
	"fmt"
	. "jean/instructions/base"
)

type InstructionFactory struct {
	codeMap map[byte]func() Instruction
}

func (factory *InstructionFactory) AddInstruction(bytecode byte, value func() Instruction) {
	factory.codeMap[bytecode] = value
}

var Factory *InstructionFactory

func init() {
	Factory = &InstructionFactory{codeMap: make(map[byte]func() Instruction, 256)}
}

func (factory *InstructionFactory) NewInstruction(opcode byte) Instruction {
	instFunc, ok := factory.codeMap[opcode]
	if !ok {
		panic(fmt.Errorf("unsupported opcode: 0x%x", opcode))
	}
	return instFunc()
}
