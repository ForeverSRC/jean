package base

import (
	"jean/rtda/jvmstack"
)

func Branch(frame *jvmstack.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
