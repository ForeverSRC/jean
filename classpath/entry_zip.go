package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var suffixList = []string{".jar", ".JAR", ".zip", ".ZIP"}

// ZipEntry 压缩包
type ZipEntry struct {
	absPath string
	zipRc   *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{
		absPath: absPath,
	}
}

func (ze *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	if ze.zipRc == nil {
		err := ze.openJar()
		if err != nil {
			return nil, nil, err
		}
	}

	classFile := ze.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}

	data, err := ze.readClassInfo(classFile)

	return data, ze, err
}

func (ze *ZipEntry) openJar() error {
	r, err := zip.OpenReader(ze.absPath)
	if err == nil {
		ze.zipRc = r
	}

	return err
}

func (ze *ZipEntry) findClass(className string) *zip.File {
	for _, f := range ze.zipRc.File {
		if f.Name == className {
			return f
		}
	}

	return nil
}

func (ze *ZipEntry) readClassInfo(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ze *ZipEntry) String() string {
	return ze.absPath
}

func needZipEntry(path string) bool {
	for _, suffix := range suffixList {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}

	return false
}
