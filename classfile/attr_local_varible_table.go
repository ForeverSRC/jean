package classfile

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableEntry
}

type LocalVariableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (lvta *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	lvta.localVariableTable = make([]*LocalVariableEntry, localVariableTableLength)
	for i := range lvta.localVariableTable {
		lvta.localVariableTable[i] = &LocalVariableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
}
