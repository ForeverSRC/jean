package instructions

import (
	"fmt"
	"jean/classfile"
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda"
)

import (
	_ "jean/instructions/comparisons"
	_ "jean/instructions/constants"
	_ "jean/instructions/control"
	_ "jean/instructions/conversions"
	_ "jean/instructions/extended"
	_ "jean/instructions/loads"
	_ "jean/instructions/math"
	_ "jean/instructions/stack"
	_ "jean/instructions/stores"
)

func Interpreter(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()

	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)

	// temp
	defer catchErr(frame)
	loop(thread, bytecode)
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		// calculate pc
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := factory.Factory.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
