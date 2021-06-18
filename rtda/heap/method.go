package heap

import "jean/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}

	return methods
}

func (m *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		m.maxStack = codeAttr.MaxStack()
		m.maxLocals = codeAttr.MaxLocals()
		m.code = codeAttr.Code()
	}
}

func (m *Method) IsSynchronized() bool {
	return m.accessFlags&ACC_SYNCHRONIZED != 0
}

func (m *Method) IsBridge() bool {
	return m.accessFlags&ACC_BRIDGE != 0
}

func (m *Method) IsVarargs() bool {
	return m.accessFlags&ACC_VARARGS != 0
}

func (m *Method) IsNative() bool {
	return m.accessFlags&ACC_NATIVE != 0
}

func (m *Method) IsAbstract() bool {
	return m.accessFlags&ACC_ABSTRACT != 0
}

func (m *Method) IsStrict() bool {
	return m.accessFlags&ACC_STRICT != 0
}

func (m *Method) MaxStack() uint {
	return m.maxStack
}
func (m *Method) MaxLocals() uint {
	return m.maxLocals
}
func (m *Method) Code() []byte {
	return m.code
}
