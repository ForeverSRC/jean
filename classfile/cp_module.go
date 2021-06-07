package classfile

type ConstantModuleInfo struct {
	nameIndex uint16
}

func (cm *ConstantModuleInfo) readInfo(reader *ClassReader) {
	cm.nameIndex = reader.readUint16()
}

type ConstantPackageInfo struct {
	nameIndex uint16
}

func (cp *ConstantPackageInfo) readInfo(reader *ClassReader) {
	cp.nameIndex = reader.readUint16()
}
