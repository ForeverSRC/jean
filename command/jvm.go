package command

import (
	"fmt"
	"jean/classpath"
	"jean/constants"
	"jean/instructions"
	"jean/instructions/base"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
	"strings"
	"time"
)

type JVM struct {
	cmd         *Cmd
	classLoader *heap.ClassLoader
	mainThread  *jvmstack.Thread
}

func NewJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)

	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	return &JVM{
		cmd:         cmd,
		classLoader: classLoader,
		mainThread:  jvmstack.NewThread(),
	}
}

func (jvm *JVM) Start() {
	jvm.initVM()
	jvm.execMain()
}

func (jvm *JVM) initVM() {
	fmt.Println("JVM start initializing....")
	startTime := time.Now()
	vmClass := jvm.classLoader.LoadClass(constants.SunMiscVM)
	base.InitClass(jvm.mainThread, vmClass)
	instructions.Interpreter(jvm.mainThread, false)
	elapsedTime := time.Since(startTime) / time.Millisecond  // duration in ms
	fmt.Printf("JVM successfully initialized. Cost: %d ms\n",elapsedTime)
}

func (jvm *JVM) execMain() {
	className := strings.Replace(jvm.cmd.class, ".", "/", -1)
	mainClass := jvm.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()

	if mainMethod == nil {
		fmt.Printf("Main method not found in class %s\n", jvm.cmd.class)
		return
	}

	argsArr := jvm.createArgsArray()
	frame := jvm.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argsArr) // 给main方法传递String[] args
	jvm.mainThread.PushFrame(frame)
	instructions.Interpreter(jvm.mainThread, jvm.cmd.verboseInstFlag)
}

// createArgsArray return String[] args
func (jvm *JVM) createArgsArray() *heap.Object {
	stringClass := jvm.classLoader.LoadClass(constants.JavaLangString)
	argsLen := uint(len(jvm.cmd.args))

	argsArr := stringClass.ArrayClass().NewArray(argsLen)
	jArgs := argsArr.Refs()
	for i, arg := range jvm.cmd.args {
		jArgs[i] = heap.JString(jvm.classLoader, arg)
	}

	return argsArr

}
