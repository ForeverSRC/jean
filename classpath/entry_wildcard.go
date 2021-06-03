package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// 通配符类路径 指定某个目录下所有jar文件
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] //remove *
	compositeEntry := make([]Entry, 0)

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	err := filepath.Walk(baseDir, walkFn)
	if err != nil {
		panic(err)
	}
	return compositeEntry
}

func needWildcardEntry(path string) bool {
	return strings.HasSuffix(path, "*")
}
