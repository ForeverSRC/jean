package command

import (
	"fmt"
	"jean/classfile"
	"jean/classpath"
	"jean/instructions"
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
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)

	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		instructions.Interpreter(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		fmt.Printf("method name: %s; descriptor: %s\n", m.Name(), m.Descriptor())
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}

	return nil
}
