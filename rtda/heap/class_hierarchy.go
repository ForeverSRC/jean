package heap

// isAssignableFrom c is assignable from other means
// 1.c and other are the same class
// 2.other is sub class of c
// 3.other implemented interface c
func (c *Class) isAssignableFrom(other *Class) bool {
	if other == c {
		return true
	}

	if !c.IsInterface() {
		return other.isSubClassOf(c)
	} else {
		return other.isImplements(c)
	}
}

// isSubClassOf 判断S是否是T的子类，实际上也就是判断T是否是S的（直接或间接）超类。
func (c *Class) isSubClassOf(other *Class) bool {
	for super := c.superClass; super != nil; super = super.superClass {
		if super == other {
			return true
		}
	}

	return false
}

// isImplements 判断S是否实现了T接口，就看S或S的（直接或间接）超类是否实现了某个接口T' ,T’要么是T，要么是T的子接口。
func (c *Class) isImplements(iface *Class) bool {
	for class := c; class != nil; class = class.superClass {
		for _, i := range class.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}

	return false
}

func (c *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range c.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}

	return false
}
