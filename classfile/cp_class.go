package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (cc *ConstantClassInfo) readInfo(reader *ClassReader) {
	cc.nameIndex = reader.readUint16()
}

func (cc *ConstantClassInfo) Name() string {
	return cc.cp.getUtf8(cc.nameIndex)
}
