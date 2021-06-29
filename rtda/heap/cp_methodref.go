package heap

import "jean/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(rtCp *ConstantPool, info *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.rtCp = rtCp
	ref.copyMemberRefInfo(&info.ConstantMemberrefInfo)
	return ref
}

func (mr *MethodRef) ResolvedMethod() *Method {
	if mr.method == nil {
		mr.resolveMethodRef()
	}

	return mr.method
}

func (mr *MethodRef) resolveMethodRef() {
	d := mr.rtCp.class
	c := mr.ResolvedClass()
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(c, mr.name, mr.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	mr.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterface(class.interfaces, name, descriptor)
	}
	return method
}
