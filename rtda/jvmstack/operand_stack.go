package jvmstack

import (
	"jean/rtda/heap"
	"math"
)

type OperandStack struct {
	top   uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}

	return nil
}

func (os *OperandStack) PushInt(val int32) {
	os.slots[os.top].num = val
	os.top++
}

func (os *OperandStack) PopInt() int32 {
	os.top--
	return os.slots[os.top].num
}

func (os *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	os.slots[os.top].num = int32(bits)
	os.top++
}

func (os *OperandStack) PopFloat() float32 {
	os.top--
	bits := uint32(os.slots[os.top].num)
	return math.Float32frombits(bits)
}

func (os *OperandStack) PushLong(val int64) {
	os.slots[os.top].num = int32(val)
	os.slots[os.top+1].num = int32(val >> 32)
	os.top += 2
}

func (os *OperandStack) PopLong() int64 {
	os.top -= 2
	low := uint32(os.slots[os.top].num)
	high := uint32(os.slots[os.top+1].num)
	return int64(high)<<32 | int64(low)
}

func (os *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	os.PushLong(int64(bits))
}

func (os *OperandStack) PopDouble() float64 {
	bits := uint64(os.PopLong())
	return math.Float64frombits(bits)
}

func (os *OperandStack) PushRef(ref *heap.Object) {
	os.slots[os.top].ref = ref
	os.top++
}

func (os *OperandStack) PopRef() *heap.Object {
	os.top--
	ref := os.slots[os.top].ref

	// help to Golang gc
	os.slots[os.top].ref = nil

	return ref
}

func (os *OperandStack) PushSlot(slot Slot) {
	os.slots[os.top] = slot
	os.top++
}

func (os *OperandStack) PopSlot() Slot {
	os.top--
	return os.slots[os.top]
}

func (os *OperandStack) TopSlot() Slot {
	return os.slots[os.top-1]
}

// GetRefFromTop 返回距离栈顶n个slot的引用变量
func (os *OperandStack) GetRefFromTop(n uint) *heap.Object {
	return os.slots[os.top-1-n].ref
}

func (os *OperandStack) PushBoolean(val bool) {
	if val {
		os.PushInt(1)
	} else {
		os.PushInt(0)
	}
}

func (os *OperandStack) PopBoolean() bool {
	return os.PopInt() == 1
}

func (os *OperandStack) Clear() {
	os.top = 0
	for i := range os.slots {
		// help gc
		os.slots[i].ref = nil
	}
}
