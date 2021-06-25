package heap

import (
	"jean/classfile"
)

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}

	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}

	return method
}

func (m *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		m.maxStack = codeAttr.MaxStack()
		m.maxLocals = codeAttr.MaxLocals()
		m.code = codeAttr.Code()
	}
}

func (m *Method) calArgSlotCount(parameterTypes []string) {
	for _, paramType := range parameterTypes {
		m.argSlotCount++
		if paramType == "J" || paramType == "D" {
			m.argSlotCount++
		}
	}
	// instant method, need a slot for param 'this'
	if !m.IsStatic() {
		m.argSlotCount++
	}
}

func (m *Method) injectCodeAttribute(returnType string) {
	m.maxStack = 4
	m.maxLocals = m.argSlotCount

	switch returnType[0] {
	case 'V':
		m.code = []byte{0xfe, 0xb1} //return
	case 'D':
		m.code = []byte{0xfe, 0xaf} //dreturn
	case 'F':
		m.code = []byte{0xfe, 0xae} //freturn
	case 'J':
		m.code = []byte{0xfe, 0xad} //lreturn
	case 'L', '[':
		m.code = []byte{0xfe, 0xb0} //areturn
	default:
		m.code = []byte{0xfe, 0xac} //ireturn
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

func (m *Method) ArgSlotCount() uint {
	return m.argSlotCount
}
