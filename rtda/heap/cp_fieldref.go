package heap

import "jean/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(rtCp *ConstantPool, info *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.rtCp = rtCp
	ref.copyMemberRefInfo(&info.ConstantMemberrefInfo)
	return ref
}

func (f *FieldRef) ResolvedField() *Field {
	if f.field == nil {
		f.resolveFieldRef()
	}
	return f.field
}

func (f *FieldRef) resolveFieldRef() {
	// d-->c.someFiled

	d := f.rtCp.class
	c := f.ResolvedClass()
	field := lookupField(c, f.name, f.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}

	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	f.field = field

}

func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}
