package references

import (
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

const (
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)

var aTypeMap = map[uint8]string{
	AT_BOOLEAN: "[Z",
	AT_CHAR:    "[C",
	AT_FLOAT:   "[F",
	AT_DOUBLE:  "[D",
	AT_BYTE:    "[B",
	AT_SHORT:   "[S",
	AT_INT:     "[I",
	AT_LONG:    "[J",
}

type NEW_ARRAY struct {
	atype uint8
}

func (na *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	na.atype = reader.ReadUint8()
}

func (na *NEW_ARRAY) Execute(frame *jvmstack.Frame) {
	stack := frame.OperandStack()
	count := stack.PopInt()

	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	classLoader := frame.Method().Class().Loader()
	arrClass := getPrimitiveArrayClass(classLoader, na.atype)
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)

}

func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	at, ok := aTypeMap[atype]
	if !ok {
		panic("Invalid atype!")
	}

	return loader.LoadClass(at)
}

func init() {
	factory.Factory.AddInstruction(0xbc, func() base.Instruction {
		return &NEW_ARRAY{}
	})
}
