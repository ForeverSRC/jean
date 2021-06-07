package classfile

type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (cmh *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	cmh.referenceKind = reader.readUint8()
	cmh.referenceIndex = reader.readUint16()
}
