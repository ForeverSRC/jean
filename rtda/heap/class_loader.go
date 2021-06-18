package heap

import (
	"fmt"
	"jean/classfile"
	"jean/classpath"
)

type ClassLoader struct {
	cp *classpath.ClassPath

	// classMap loaded class
	classMap map[string]*Class
}

func NewClassLoader(cp *classpath.ClassPath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

func (cl *ClassLoader) LoadClass(name string) *Class {
	if class, ok := cl.classMap[name]; ok {
		// 类已经加载
		return class
	}

	return cl.LoadNonArrayClass(name)
}

func (cl *ClassLoader) LoadNonArrayClass(name string) *Class {
	data, entry := cl.readClass(name)
	class := cl.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

func (cl *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := cl.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}

	return data, entry
}

func (cl *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = cl
	resolveSuperClass(class)
	resolveInterfaces(class)
	cl.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}

	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	// TODO Java虚拟机规范4.10节详细介绍了类的验证算法，留待以后视情况实现
}

func prepare(class *Class) {
	calInstanceFieldSlotIds(class)
	calStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func calInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	// 父类在子类之前完成加载
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}

	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}

	class.instanceSlotCount = slotId
}

func calStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	// 注意：静态变量属于类

	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.IsLongOrDouble() {
				slotId++
			}
		}
	}

	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = NewSlots(class.staticSlotCount)
	for _, field := range class.fields {
		// static final
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	rtCp := class.constantPool
	cpIndex := field.ConstantValueIndex()
	slotId := field.slotId

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := rtCp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := rtCp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := rtCp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := rtCp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			// todo:implements
			panic("todo")
		}
	}
}
