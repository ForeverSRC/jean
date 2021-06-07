package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (cnt *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	cnt.nameIndex = reader.readUint16()
	cnt.descriptorIndex = reader.readUint16()
}
