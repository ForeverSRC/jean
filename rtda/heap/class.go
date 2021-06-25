package heap

import (
	"jean/classfile"
	"jean/constants"
	"strings"
)

type Class struct {
	accessFlag        uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
	// java.lang.Class 实例
	jClass *Object
}

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlag = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (c *Class) NewObject() *Object {
	return newObject(c)
}

func (c *Class) IsPublic() bool {
	return c.accessFlag&ACC_PUBLIC != 0
}

func (c *Class) IsProtected() bool {
	return c.accessFlag&ACC_PROTECTED != 0
}

func (c *Class) IsFinal() bool {
	return c.accessFlag&ACC_FINAL != 0
}

func (c *Class) IsSuper() bool {
	return c.accessFlag&ACC_SUPER != 0
}

func (c *Class) IsInterface() bool {
	return c.accessFlag&ACC_INTERFACE != 0
}

func (c *Class) IsAbstract() bool {
	return c.accessFlag&ACC_ABSTRACT != 0
}

func (c *Class) IsAnnotation() bool {
	return c.accessFlag&ACC_ANNOTATION != 0
}

func (c *Class) IsEnum() bool {
	return c.accessFlag&ACC_ENUM != 0
}

// IsAccessibleTo c能否被other访问
func (c *Class) IsAccessibleTo(other *Class) bool {
	return c.IsPublic() || c.GetPackageName() == other.GetPackageName()
}

func (c *Class) IsJlObject() bool {
	return c.name == constants.JavaLangObject
}

func (c *Class) IsJlCloneable() bool {
	return c.name == constants.JavaLangCloneable
}

func (c *Class) IsJioSerializable() bool {
	return c.name == constants.JavaIoSerializable
}

// GetPackageName exp: "java/lang/Integer" packageName is "java/lang"
func (c *Class) GetPackageName() string {
	if i := strings.LastIndex(c.name, "/"); i >= 0 {
		return c.name[:i]
	}

	return ""
}

func (c *Class) ConstantPool() *ConstantPool {
	return c.constantPool
}

func (c *Class) StaticVars() Slots {
	return c.staticVars
}

func (c *Class) SuperClass() *Class {
	return c.superClass
}

func (c *Class) Name() string {
	return c.name
}

func (c *Class) Loader() *ClassLoader {
	return c.loader
}

func (c *Class) JClass() *Object {
	return c.jClass
}

func (c *Class) InitStarted() bool {
	return c.initStarted
}

func (c *Class) StartInit() {
	c.initStarted = true
}

func (c *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(c.name)
	return c.loader.LoadClass(arrayClassName)
}

func (c *Class) getField(name, descriptor string, isStatic bool) *Field {
	for cl := c; cl != nil; cl = cl.superClass {
		for _, field := range cl.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}

	return nil
}

func (c *Class) JavaName() string {
	return strings.Replace(c.name, "/", ".", -1)
}

func (c *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[c.name]
	return ok
}

func (c *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := c.getField(fieldName, fieldDescriptor, true)
	return c.staticVars.GetRef(field.slotId)
}

func (c *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := c.getField(fieldName, fieldDescriptor, true)
	c.staticVars.SetRef(field.slotId, ref)
}

func (c *Class) GetClinitMethod() *Method {
	return c.getMethod("<clinit>", "()V", true)
}

func (c *Class) GetInstanceMethod(name, descriptor string) *Method {
	return c.getMethod(name, descriptor, false)
}

func (c *Class) GetMainMethod() *Method {
	return c.getMethod("main", "([Ljava/lang/String;)V", true)
}

func (c *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for clazz := c; clazz != nil; clazz = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {
				return method
			}
		}
	}

	return nil
}
