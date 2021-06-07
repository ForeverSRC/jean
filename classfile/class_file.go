package classfile

import "fmt"

const leftBoundOfMajorVersion = 45
const rightBoundOfMajorVersion = 52
const classFileMagic = 0xCAFEBABE

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	cf.superClass = reader.readUint16()
	cf.interfaces = reader.readUint16s()
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)
}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != classFileMagic {
		// todo 后面改成异常抛出
		panic("java.lang.ClassFormatError: magic! ")
	}
	cf.magic = magic
}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()

	if cf.majorVersion == leftBoundOfMajorVersion {
		return
	}

	if leftBoundOfMajorVersion < cf.majorVersion && cf.majorVersion <= rightBoundOfMajorVersion {
		if cf.minorVersion == 0 {
			return
		}
	}

	// todo 后面改成异常抛出
	panic("java.lang.UnsupportClassVersionError! ")
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}

	// 只有java.lang.Object没有父类
	return ""
}

func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		interfaceNames[i] = cf.constantPool.getClassName(cpIndex)
	}

	return interfaceNames
}

func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}
