package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (cs *ConstantStringInfo) readInfo(reader *ClassReader) {
	cs.stringIndex = reader.readUint16()
}

func (cs *ConstantStringInfo) String() string {
	return cs.cp.getUtf8(cs.stringIndex)
}
