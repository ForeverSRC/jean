package instructions

import (
	"fmt"
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

import (
	_ "jean/instructions/comparisons"
	_ "jean/instructions/constants"
	_ "jean/instructions/control"
	_ "jean/instructions/conversions"
	_ "jean/instructions/extended"
	_ "jean/instructions/loads"
	_ "jean/instructions/math"
	_ "jean/instructions/references"
	_ "jean/instructions/stack"
	_ "jean/instructions/stores"
)

func Interpreter(method *heap.Method) {
	thread := jvmstack.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	// temp
	defer catchErr(frame)
	loop(thread, method.Code())
}

func loop(thread *jvmstack.Thread, bytecode []byte) {
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

func catchErr(frame *jvmstack.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}
