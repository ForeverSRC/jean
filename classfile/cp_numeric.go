package classfile

import (
	"math"
)

type ConstantIntegerInfo struct {
	val int32
}

func (ci *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	ci.val = int32(bytes)
}

func (ci *ConstantIntegerInfo) Value() int32 {
	return ci.val
}

type ConstantFloatInfo struct {
	val float32
}

func (cf *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	cf.val = math.Float32frombits(bytes)
}

func (cf *ConstantFloatInfo) Value() float32 {
	return cf.val
}

type ConstantLongInfo struct {
	val int64
}

func (cl *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	cl.val = int64(bytes)
}

func (cl *ConstantLongInfo) Value() int64 {
	return cl.val
}

type ConstantDoubleInfo struct {
	val float64
}

func (cd *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	cd.val = math.Float64frombits(bytes)
}

func (cd *ConstantDoubleInfo) Value() float64 {
	return cd.val
}
