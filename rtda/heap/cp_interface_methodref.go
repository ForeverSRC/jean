package heap

import "jean/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(rtCp *ConstantPool, info *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.rtCp = rtCp
	ref.copyMemberRefInfo(&info.ConstantMemberrefInfo)
	return ref
}

func (imr *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if imr.method == nil {
		imr.resolveInterfaceMethodRef()
	}

	return imr.method
}

func (imr *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d := imr.rtCp.class
	c := imr.ResolvedClass()

	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupInterfaceMethod(c, imr.name, imr.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	imr.method = method

}

func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}

	return lookupMethodInInterface(iface.interfaces, name, descriptor)
}
