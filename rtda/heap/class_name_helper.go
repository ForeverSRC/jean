package heap

// class name ->descriptor
var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

// [XXX -> [[XXX
// int -> [I
// XXX -> [LXXX;
func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

func toDescriptor(className string) string {
	if className[0] == '[' {
		return className
	}

	if d, ok := primitiveTypes[className]; ok {
		return d
	}

	return "L" + className + ";"
}

func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}

	panic("Not array: " + className)
}

func toClassName(descriptor string) string {
	// 数组，类名就是描述符名
	if descriptor[0] == '[' {
		return descriptor
	}

	// 引用，例如Ljava/lang/String; 返回java/lang/String
	if descriptor[0] == 'L' {
		return descriptor[1 : len(descriptor)-1]
	}

	// 基本类型
	for className, d := range primitiveTypes {
		if d == descriptor {
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}
