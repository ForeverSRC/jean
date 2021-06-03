package command

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	// 帮助位
	helpFlag bool

	// 版本位
	versionFlag bool

	// classpath
	cpOption string

	// -Xjre
	XjreOption string

	class string
	args  []string
}

func ParseCmd() *Cmd {
	cmd := Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "? ", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")

	// 解析选项
	flag.Parse()

	// 捕获其他未被解析的参数
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return &cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
