package heap

import (
	"fmt"
	"jean/classfile"
	"jean/constants"
	"strings"
)

type Class struct {
	accessFlag     uint16
	name           string
	superClassName string
	interfaceNames []string
	constantPool   *ConstantPool
	fields         []*Field
	// 本类中的方法
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
	// java.lang.Class 实例
	jClass     *Object
	sourceFile string

	// vtable name&&descriptor -> Method
	// 包含父类的方法
	vtable map[string]*Method

	itable map[string]*Method
}

const tableKey = "%s %s"

func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{
		vtable: make(map[string]*Method),
		itable: make(map[string]*Method),
	}
	class.accessFlag = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	class.sourceFile = getSourceFile(cf)
	return class
}

func getSourceFile(cf *classfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}

	return "Unknown"
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
func (c *Class) AccessFlags() uint16 {
	return c.accessFlag
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

func (c *Class) SourceFile() string {
	return c.sourceFile
}

func (c *Class) InitStarted() bool {
	return c.initStarted
}

func (c *Class) StartInit() {
	c.initStarted = true
}

func (c *Class) Interfaces() []*Class {
	return c.interfaces
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

func (c *Class) GetStaticMethod(name, descriptor string) *Method {
	return c.getMethod(name, descriptor, true)
}

func (c *Class) GetConstructor(descriptor string) *Method {
	return c.GetInstanceMethod("<init>", descriptor)
}

func (c *Class) GetMainMethod() *Method {
	return c.getMethod("main", "([Ljava/lang/String;)V", true)
}

func (c *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for clazz := c; clazz != nil; clazz = clazz.superClass {
		for _, method := range clazz.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {
				return method
			}
		}
	}

	return nil
}

func (c *Class) GetFields(publicOnly bool) []*Field {
	if publicOnly {
		publicFields := make([]*Field, 0, len(c.fields))
		for _, field := range c.fields {
			if field.IsPublic() {
				publicFields = append(publicFields, field)
			}
		}
		return publicFields
	} else {
		return c.fields
	}
}

func (c *Class) GetMethods(publicOnly bool) []*Method {
	methods := make([]*Method, 0, len(c.methods))
	for _, method := range c.methods {
		if !method.isClinit() && !method.isConstructor() {
			if !publicOnly || method.IsPublic() {
				methods = append(methods, method)
			}
		}
	}
	return methods
}

func (c *Class) GetConstructors(publicOnly bool) []*Method {
	constructors := make([]*Method, 0, len(c.methods))
	for _, method := range c.methods {
		if method.isConstructor() {
			if !publicOnly || method.IsPublic() {
				constructors = append(constructors, method)
			}
		}
	}
	return constructors
}

func (c *Class) SetVtable(name, descriptor string, method *Method) {
	c.vtable[fmt.Sprintf(tableKey, name, descriptor)] = method
}

func (c *Class) GetFromVtable(name, descriptor string) *Method {
	m, ok := c.vtable[fmt.Sprintf(tableKey, name, descriptor)]
	if !ok {
		return nil
	}

	return m
}

func (c *Class) SetItable(name, descriptor string, method *Method) {
	c.itable[fmt.Sprintf(tableKey, name, descriptor)] = method
}

func (c *Class) GetFromItable(name, descriptor string) *Method {
	m, ok := c.itable[fmt.Sprintf(tableKey, name, descriptor)]
	if !ok {
		return nil
	}

	return m
}
