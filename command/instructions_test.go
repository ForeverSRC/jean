package command

import (
	"fmt"
	"jean/classpath"
	"jean/instructions"
	"jean/rtda/heap"
	"strings"
	"testing"
)

func TestInstructions(t *testing.T) {
	cmd := ParseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startTestInstructions(cmd)
	}
}

func startTestInstructions(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
	classLoader:=heap.NewClassLoader(cp)
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass:=classLoader.LoadClass(className)

	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		instructions.Interpreter(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
