package main

import "jean/command"

func main() {
	cmd := command.ParseCmd()
	ok := cmd.ParseFlags()
	if ok {
		command.NewJVM(cmd).Start()
	}
}
