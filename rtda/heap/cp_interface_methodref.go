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
