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
