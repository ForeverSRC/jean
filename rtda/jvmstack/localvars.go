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

func (s LocalVars) SetInt(index uint, val int32) {
	s[index].SetNum(val)
}

func (s LocalVars) GetInt(index uint) int32 {
	return s[index].Num()
}

func (s LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	s[index].SetNum(int32(bits))
}

func (s LocalVars) GetFloat(index uint) float32 {
	bits := uint32(s[index].Num())
	return math.Float32frombits(bits)
}

func (s LocalVars) SetLong(index uint, val int64) {
	s[index].SetNum(int32(val))
	s[index+1].SetNum(int32(val >> 32))
}

func (s LocalVars) GetLong(index uint) int64 {
	low := uint32(s[index].Num())
	high := uint32(s[index+1].Num())
	return int64(high)<<32 | int64(low)
}

func (s LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	s.SetLong(index, int64(bits))
}

func (s LocalVars) GetDouble(index uint) float64 {
	bits := uint64(s.GetLong(index))
	return math.Float64frombits(bits)
}

func (s LocalVars) SetRef(index uint, ref *heap.Object) {
	s[index].SetRef(ref)
}

func (s LocalVars) GetRef(index uint) *heap.Object {
	return s[index].Ref()
}
