package classpath

import (
	"errors"
	"strings"
)

// CompositeEntry 复合目录类路径
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	paths := strings.Split(pathList, pathListSeparator)
	compositeEntry := make([]Entry, len(paths))
	for idx, path := range paths {
		entry := newEntry(path)
		compositeEntry[idx] = entry
	}

	return compositeEntry
}

func (ce CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range ce {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (ce CompositeEntry) String() string {
	strs := make([]string, len(ce))
	for i, entry := range ce {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}

func needCompositeEntry(path string) bool {
	return strings.Contains(path, pathListSeparator)
}
