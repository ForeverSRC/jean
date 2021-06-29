package io

import (
	"jean/constants"
	"jean/native"
	"jean/rtda/jvmstack"
	"os"
	"unsafe"
)

// private static native void initIDs();
func fosInitIDs(frame *jvmstack.Frame) {
	// todo
}

// private native void writeBytes(byte b[],int off,int len,boolean append) throws IoException
func writeBytes(frame *jvmstack.Frame) {
	vars := frame.LocalVars()
	b := vars.GetRef(1)
	off := vars.GetInt(2)
	length := vars.GetInt(3)

	jBytes := b.Data().([]int8)
	goBytes := castInt8ToUint8s(jBytes)
	goBytes = goBytes[off : off+length]
	os.Stdout.Write(goBytes)
}

func castInt8ToUint8s(jBytes []int8) (goBytes []uint8) {
	ptr := unsafe.Pointer(&jBytes)
	goBytes = *((*[]byte)(ptr))
	return
}

func init() {
	native.Register(constants.JavaIoFileOutputStream, "writeBytes", "([BIIZ)V", writeBytes)
	native.Register(constants.JavaIoFileOutputStream, "initIDs", "()V", fosInitIDs)
}
