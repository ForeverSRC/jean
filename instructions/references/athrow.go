package references

import (
	"fmt"
	"jean/instructions/base"
	"jean/instructions/factory"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
	"reflect"
)

type ATHROW struct {
	base.NoOperandsInstruction
}

func (at *ATHROW) Execute(frame *jvmstack.Frame) {
	ex := frame.OperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}

	thread := frame.Thread()
	if !findAndGotoExceptionHandler(thread, ex) {
		handleUncaughtException(thread, ex)
	}
}

func findAndGotoExceptionHandler(thread *jvmstack.Thread, ex *heap.Object) bool {
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC() - 1

		handlerPC := frame.Method().FindExceptionHandler(ex.Class(), pc)
		if handlerPC > 0 {
			stack := frame.OperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handlerPC)
			return true
		}

		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}

	return false
}

func handleUncaughtException(thread *jvmstack.Thread, ex *heap.Object) {
	thread.ClearStack()

	jMsg := ex.GetRefVar("detailMessage", "Ljava/lang/String;")
	goMsg := heap.GoString(jMsg)
	fmt.Printf("%s: %s", ex.Class().JavaName(), goMsg)

	// java虚拟机栈清空，解释器终止执行，采用go语言打印Java虚拟机栈信息
	stes := reflect.ValueOf(ex.Extra())
	for i := 0; i < stes.Len(); i++ {
		ste := stes.Index(i).Interface().(interface {
			String() string
		})
		fmt.Printf("\tat %s", ste.String())
	}
}

func init() {
	athrow := &ATHROW{}

	factory.Factory.AddInstruction(0xbf, func() base.Instruction {
		return athrow
	})
}
