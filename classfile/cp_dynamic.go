package classfile

type ConstantDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (cd *ConstantDynamicInfo) readInfo(reader *ClassReader) {
	cd.bootstrapMethodAttrIndex = reader.readUint16()
	cd.nameAndTypeIndex = reader.readUint16()
}

type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (cid *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	cid.bootstrapMethodAttrIndex = reader.readUint16()
	cid.nameAndTypeIndex = reader.readUint16()
}
