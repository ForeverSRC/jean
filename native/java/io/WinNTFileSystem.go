package io

import (
	"jean/native"
	"jean/rtda/jvmstack"
)

const winNtFs = "java/io/WinNTFileSystem"

func winfsInitIDs(frame *jvmstack.Frame) {
	// todo
}

func init() {
	native.Register(winNtFs, "initIDs", "()V", winfsInitIDs)
}
