package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (lnta LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	lnta.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range lnta.lineNumberTable {
		lnta.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

func (lnta *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(lnta.lineNumberTable) - 1; i >= 0; i-- {
		entry := lnta.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}

	return -1
}
