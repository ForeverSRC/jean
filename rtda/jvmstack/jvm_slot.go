package jvmstack

import "jean/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}

func (s *Slot) Num() int32 {
	return s.num
}

func (s *Slot) SetNum(num int32) {
	s.num = num
}

func (s *Slot) Ref() *heap.Object {
	return s.ref
}

func (s *Slot) SetRef(obj *heap.Object) {
	s.ref = obj
}
