package heap

import "jean/classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}

	return fields
}

func (f *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		f.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (f *Field) IsVolatile() bool {
	return f.accessFlags&ACC_VOLATILE != 0
}

func (f *Field) IsTransient() bool {
	return f.accessFlags&ACC_TRANSIENT != 0
}

func (f *Field) IsEnum() bool {
	return f.accessFlags&ACC_ENUM != 0
}

func (f *Field) IsLongOrDouble() bool {
	return f.descriptor == "J" || f.descriptor == "D"
}

func (f *Field) ConstantValueIndex() uint {
	return f.constValueIndex
}

func (f *Field) SlotId() uint {
	return f.slotId
}

// reflection
func (f *Field) Type() *Class {
	className := toClassName(f.descriptor)
	return f.class.loader.LoadClass(className)
}
