package heap

import (
	"jean/classfile"
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
	return c.IsPublic() || c.getPackageName() == other.getPackageName()
}

// exp: "java/lang/Integer" packageName is "java/lang"
func (c *Class) getPackageName() string {
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

func (c *Class) GetMainMethod() *Method {
	return c.getStaticMethod("main","([Ljava/lang/String;)V")
}

func (c *Class) getStaticMethod(name,descriptor string) *Method {
	for _,method:=range c.methods{
		if method.IsStatic()&&method.name==name&&method.descriptor==descriptor{
			return method
		}
	}

	return nil
}