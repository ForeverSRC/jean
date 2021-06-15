package rtda

import "math"

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

func (os *OperandStack) PushRef(ref *Object) {
	os.slots[os.top].ref = ref
	os.top++
}

func (os *OperandStack) PopRef() *Object {
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
