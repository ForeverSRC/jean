package lang

import (
	"jean/constants"
	"jean/native"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

// private native Throwable fillInStackTrace(int dummy)
func fillInStackTrace(frame *jvmstack.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)

	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

func createStackTraceElements(tObj *heap.Object, thread *jvmstack.Thread) []*StackTraceElement {
	// 由于栈顶两帧正在执行private native Throwable fillInStackTrace(int dummy)
	// 和 public synchronized Throwable fillInStackTrace() 方法，所以需要跳过这两帧 (+2)
	// 这两帧下面的几帧正在执行异常类的构造函数，所以也要跳过，具体要跳过多少帧数则要看异常类的继承层次。
	// distanceToObject()函数计算所需跳过的帧数
	skip := distanceToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]

	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}

	return stes
}

func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}

	return distance
}

func createStackTraceElement(frame *jvmstack.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName:   class.SourceFile(),
		className:  class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}

func init() {
	native.Registrer(constants.JavaLangThrowable, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}
