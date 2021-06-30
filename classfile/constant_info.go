package classfile

import "fmt"

const (
	CONSTANT_Utf8               = 1
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_Class              = 7
	CONSTANT_String             = 8
	CONSTANT_FieldRef           = 9
	CONSTANT_MethodRef          = 10
	CONSTANT_InterfaceMethodRef = 11
	CONSTANT_NameAndType        = 12
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_Dynamic            = 17
	CONSTANT_InvokeDynamic      = 18
	CONSTANT_Moudule            = 19
	CONSTANT_Package            = 20
)

type ConstantInfo interface {
	// readInfo 读取常量信息，需要由具体常量结构体实现
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_FieldRef:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_MethodRef:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodRef:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_Dynamic:
		return &ConstantDynamicInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	case CONSTANT_Moudule:
		return &ConstantModuleInfo{}
	case CONSTANT_Package:
		return &ConstantPackageInfo{}
	default:
		// todo 后面改成抛异常
		panic(fmt.Errorf("java.lang.ClassFormatError: constant pool tag: %v !  ", tag))
	}
}
