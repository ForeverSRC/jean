package heap

import (
	"jean/classfile"
)

type Method struct {
	ClassMember
	maxStack                uint
	maxLocals               uint
	code                    []byte
	argSlotCount            uint
	exceptionTable          ExceptionTable // todo: rename
	lineNumberTable         *classfile.LineNumberTableAttribute
	exceptions              *classfile.ExceptionsAttribute // todo: rename
	parameterAnnotationData []byte                         // RuntimeVisibleParameterAnnotations_attribute
	annotationDefaultData   []byte                         // AnnotationDefault_attribute
	parsedDescriptor        *MethodDescriptor
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
	method.parsedDescriptor = md
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
		m.lineNumberTable = codeAttr.LineNumberTableAttribute()
		m.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(), m.class.constantPool)
	}

	m.exceptions = cfMethod.ExceptionsAttribute()
	m.annotationData = cfMethod.RuntimeVisibleAnnotationsAttributeData()
	m.parameterAnnotationData = cfMethod.RuntimeVisibleParameterAnnotationsAttributeData()
	m.annotationDefaultData = cfMethod.AnnotationDefaultAttributeData()
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

// injectCodeAttribute for native method
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

func (m *Method) FindExceptionHandler(exclass *Class, pc int) int {
	handler := m.exceptionTable.findExceptionHandler(exclass, pc)
	if handler != nil {
		return handler.handlerPc
	}

	return -1
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

func (m *Method) isConstructor() bool {
	return !m.IsStatic() && m.name == "<init>"
}
func (m *Method) isClinit() bool {
	return m.IsStatic() && m.name == "<clinit>"
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

func (m *Method) ParameterAnnotationData() []byte {
	return m.parameterAnnotationData
}
func (m *Method) AnnotationDefaultData() []byte {
	return m.annotationDefaultData
}
func (m *Method) ParsedDescriptor() *MethodDescriptor {
	return m.parsedDescriptor
}

func (m *Method) ArgSlotCount() uint {
	return m.argSlotCount
}

func (m *Method) GetLineNumber(pc int) int {
	if m.IsNative() {
		return -2
	}

	if m.lineNumberTable == nil {
		return -1
	}

	return m.lineNumberTable.GetLineNumber(pc)
}

// reflection
func (m *Method) ParameterTypes() []*Class {
	if m.argSlotCount == 0 {
		return nil
	}

	paramTypes := m.parsedDescriptor.parameterTypes
	paramClasses := make([]*Class, len(paramTypes))
	for i, paramType := range paramTypes {
		paramClassName := toClassName(paramType)
		paramClasses[i] = m.class.loader.LoadClass(paramClassName)
	}

	return paramClasses
}
func (m *Method) ReturnType() *Class {
	returnType := m.parsedDescriptor.returnType
	returnClassName := toClassName(returnType)
	return m.class.loader.LoadClass(returnClassName)
}
func (m *Method) ExceptionTypes() []*Class {
	if m.exceptions == nil {
		return nil
	}

	exIndexTable := m.exceptions.ExceptionIndexTable()
	exClasses := make([]*Class, len(exIndexTable))
	cp := m.class.constantPool

	for i, exIndex := range exIndexTable {
		classRef := cp.GetConstant(uint(exIndex)).(*ClassRef)
		exClasses[i] = classRef.ResolvedClass()
	}

	return exClasses
}
