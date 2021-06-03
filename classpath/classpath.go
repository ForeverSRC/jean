package classpath

import (
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClassPath Entry
	extClassPath  Entry
	userClassPath Entry
}

// Parse 使用-Xjre 解析启动类路径和扩展类路径 使用-classpath/-cp解析用户类路径
func Parse(jreOption, cpOption string) *ClassPath {
	cp := &ClassPath{}
	cp.parseBootAndExtClassPath(jreOption)
	cp.parseUserClasspath(cpOption)

	return cp
}

func (cp *ClassPath) parseBootAndExtClassPath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	cp.bootClassPath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreLibExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	cp.extClassPath = newWildcardEntry(jreLibExtPath)
}

func getJreDir(jreOption string) string {
	// 搜索过程：用户指定->当前目录->环境变量
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Cannot find jre folder")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func (cp *ClassPath) parseUserClasspath(cpOption string) {
	// 搜索过程：-classpath/-cp -> CLASSPATH -> 默认值 '.'
	if cpOption == "" {
		if cpe := os.Getenv("CLASSPATH"); cpe != "" {
			cpOption = cpe
		} else {
			cpOption = "."
		}
	}

	cp.userClassPath = newEntry(cpOption)
}

func (cp *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := cp.bootClassPath.readClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := cp.extClassPath.readClass(className); err == nil {
		return data, entry, err
	}

	return cp.userClassPath.readClass(className)
}

func (cp *ClassPath) String() string {
	return cp.userClassPath.String()
}
