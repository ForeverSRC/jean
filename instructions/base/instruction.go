package base

import "jean/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct{}

func (noi *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

// BranchInstruction 抽象跳转指令
type BranchInstruction struct {
	// Offset 跳转偏移量
	Offset int
}

func (bi *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	bi.Offset = int(reader.ReadInt16())
}

// Index8Instruction 抽象索引指令，索引包括局部变量表的索引
type Index8Instruction struct {
	Index uint
}

func (i8i *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	i8i.Index = uint(reader.ReadUint8())
}

// Index16Instruction 抽象常量池索引指令
type Index16Instruction struct {
	Index uint
}

func (i16i *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i16i.Index = uint(reader.ReadUint16())
}
