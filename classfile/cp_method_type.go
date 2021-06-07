package classfile

type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (cmt *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	cmt.descriptorIndex = reader.readUint16()
}
