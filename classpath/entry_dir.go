package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry 目录类路径
type DirEntry struct {
	// absDir 存放目录的绝对路径
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &DirEntry{
		absDir: absDir,
	}
}

func (de *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(de.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, de, err
}

func (de *DirEntry) String() string {
	return de.absDir
}

func DirEntryRule(path string) bool {
	return true
}
