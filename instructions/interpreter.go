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

func Interpreter(method *heap.Method, logInst bool) {
	thread := jvmstack.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	// temp
	defer catchErr(thread)
	loop(thread, logInst)
}

func loop(thread *jvmstack.Thread, logInst bool) {
	reader := &base.BytecodeReader{}

	for {
		frame := thread.CurrentFrame()
		// calculate pc
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := factory.Factory.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		// execute
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func catchErr(thread *jvmstack.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logInstruction(frame *jvmstack.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *jvmstack.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
