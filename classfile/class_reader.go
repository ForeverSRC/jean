package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

// readUint8 读取u1类型数据
func (cr *ClassReader) readUint8() uint8 {
	val := cr.data[0]
	cr.data = cr.data[1:]
	return val
}

// readUint16 读取u2类型数据
func (cr *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return val
}

// readUint32 读取u4类型数据
func (cr *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return val
}

// readUint64 Java虚拟机规范中没有规定u8
func (cr *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return val
}

// readUint16s 读取uint16表 表的大小由开头的uint16数据指出
func (cr *ClassReader) readUint16s() []uint16 {
	n := cr.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = cr.readUint16()
	}

	return s
}

// readBytes 用于读取指定数量的字节
func (cr *ClassReader) readBytes(n uint32) []byte {
	bytes := cr.data[:n]
	cr.data = cr.data[n:]
	return bytes
}
