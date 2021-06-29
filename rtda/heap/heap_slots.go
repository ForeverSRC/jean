package heap

import (
	"math"
)

type Slot struct {
	num int32
	ref *Object
}

func (s *Slot) Num() int32 {
	return s.num
}

func (s *Slot) SetNum(num int32) {
	s.num = num
}

func (s *Slot) Ref() *Object {
	return s.ref
}

func (s *Slot) SetRef(obj *Object) {
	s.ref = obj
}

type Slots []Slot

func NewSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

func (s Slots) SetInt(index uint, val int32) {
	s[index].SetNum(val)
}

func (s Slots) GetInt(index uint) int32 {
	return s[index].Num()
}

func (s Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	s[index].SetNum(int32(bits))
}

func (s Slots) GetFloat(index uint) float32 {
	bits := uint32(s[index].Num())
	return math.Float32frombits(bits)
}

func (s Slots) SetLong(index uint, val int64) {
	s[index].SetNum(int32(val))
	s[index+1].SetNum(int32(val >> 32))
}

func (s Slots) GetLong(index uint) int64 {
	low := uint32(s[index].Num())
	high := uint32(s[index+1].Num())
	return int64(high)<<32 | int64(low)
}

func (s Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	s.SetLong(index, int64(bits))
}

func (s Slots) GetDouble(index uint) float64 {
	bits := uint64(s.GetLong(index))
	return math.Float64frombits(bits)
}

func (s Slots) SetRef(index uint, ref *Object) {
	s[index].SetRef(ref)
}

func (s Slots) GetRef(index uint) *Object {
	return s[index].Ref()
}
