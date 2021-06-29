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
		switch cpInfoType := cpInfo.(type) {
		case *classfile.ConstantIntegerInfo:
			pool[i] = cpInfoType.Value()
		case *classfile.ConstantFloatInfo:
			pool[i] = cpInfoType.Value()
		case *classfile.ConstantLongInfo:
			// long 类型常量占据两个位置
			pool[i] = cpInfoType.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			// double 类型常量占据两个位置
			pool[i] = cpInfoType.Value()
			i++
		case *classfile.ConstantStringInfo:
			pool[i] = cpInfoType.String()
		case *classfile.ConstantClassInfo:
			pool[i] = newClassRef(rtCp, cpInfoType)
		case *classfile.ConstantFieldrefInfo:
			pool[i] = newFieldRef(rtCp, cpInfoType)
		case *classfile.ConstantMethodrefInfo:
			pool[i] = newMethodRef(rtCp, cpInfoType)
		case *classfile.ConstantInterfaceMethodrefInfo:
			pool[i] = newInterfaceMethodRef(rtCp, cpInfoType)
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
