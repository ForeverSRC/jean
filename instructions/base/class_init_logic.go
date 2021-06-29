package base

import (
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
)

func InitClass(thread *jvmstack.Thread, class *heap.Class) {
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)

}

func scheduleClinit(thread *jvmstack.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *jvmstack.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}
