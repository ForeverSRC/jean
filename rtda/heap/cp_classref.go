package heap

import "jean/classfile"

type ClassRef struct {
	SymbolicRef
}

func newClassRef(rtCp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.rtCp = rtCp
	ref.className = classInfo.Name()
	return ref

}
