package jvmstack

import (
	"jean/rtda/heap"
	"math"
)

type LocalVars []Slot

func NewLocalVars(slotCount uint) LocalVars {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

func (lv LocalVars) SetInt(index uint, val int32) {
	lv[index].SetNum(val)
}

func (lv LocalVars) GetInt(index uint) int32 {
	return lv[index].Num()
}

func (lv LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	lv[index].SetNum(int32(bits))
}

func (lv LocalVars) GetFloat(index uint) float32 {
	bits := uint32(lv[index].Num())
	return math.Float32frombits(bits)
}

func (lv LocalVars) SetLong(index uint, val int64) {
	lv[index].SetNum(int32(val))
	lv[index+1].SetNum(int32(val >> 32))
}

func (lv LocalVars) GetLong(index uint) int64 {
	low := uint32(lv[index].Num())
	high := uint32(lv[index+1].Num())
	return int64(high)<<32 | int64(low)
}

func (lv LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	lv.SetLong(index, int64(bits))
}

func (lv LocalVars) GetDouble(index uint) float64 {
	bits := uint64(lv.GetLong(index))
	return math.Float64frombits(bits)
}

func (lv LocalVars) SetRef(index uint, ref *heap.Object) {
	lv[index].SetRef(ref)
}

func (lv LocalVars) GetRef(index uint) *heap.Object {
	return lv[index].Ref()
}

func (lv LocalVars) SetSlot(index uint, slot Slot) {
	lv[index] = slot
}

func (lv LocalVars) GetThis() *heap.Object {
	return lv.GetRef(0)
}

func (lv LocalVars) GetBoolean(index uint) bool {
	return lv.GetInt(index) == 1
}
