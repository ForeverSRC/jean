package heap

import (
	"fmt"
	"jean/classfile"
)

type Constant interface {
}

type ConstantPool struct {
	class *Class
	pool  []Constant
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	pool := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, pool}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			pool[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			pool[i] = floatInfo
		case *classfile.ConstantLongInfo:
			// long 类型常量占据两个位置
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			pool[i] = longInfo
			i++
		case *classfile.ConstantDoubleInfo:
			// double 类型常量占据两个位置
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			pool[i] = doubleInfo
			i++
		case *classfile.ConstantStringInfo:
			stringInfo := cpInfo.(*classfile.ConstantStringInfo)
			pool[i] = stringInfo
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			pool[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldrefInfo:
			fieldRefInfo := cpInfo.(*classfile.ConstantFieldrefInfo)
			pool[i] = newFieldRef(rtCp, fieldRefInfo)
		case *classfile.ConstantMethodrefInfo:
			methodRefInfo := cpInfo.(*classfile.ConstantMethodrefInfo)
			pool[i] = newMethodRef(rtCp, methodRefInfo)
		case *classfile.ConstantInterfaceMethodrefInfo:
			interfaceMethodInfo := cpInfo.(*classfile.ConstantInterfaceMethodrefInfo)
			pool[i] = newInterfaceMethodRef(rtCp, interfaceMethodInfo)
		}

	}

	return rtCp
}

func (cp ConstantPool) GetConstant(index uint) Constant {
	if c := cp.pool[index]; c != nil {
		return c
	}

	panic(fmt.Sprintf("No constants at index %d", index))
}
