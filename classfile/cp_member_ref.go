package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (cmr *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	cmr.classIndex = reader.readUint16()
	cmr.nameAndTypeIndex = reader.readUint16()
}

func (cmr *ConstantMemberrefInfo) ClassName() string {
	return cmr.cp.getClassName(cmr.classIndex)
}

func (cmr *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return cmr.cp.getNameAndType(cmr.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
