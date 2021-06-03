package classpath

import (
	"os"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// readClass 寻找和加载class文件
	// className 为class文件的相对路径 路径之间用'/'分隔 文件名后缀 .class
	readClass(className string) ([]byte, Entry, error)

	// String 用于返回变量的字符串表示
	String() string
}

func newEntry(path string) Entry {
	if needCompositeEntry(path) {
		return newCompositeEntry(path)
	}

	if needWildcardEntry(path) {
		return newWildcardEntry(path)
	}

	if needZipEntry(path) {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
