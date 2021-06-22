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
		return other.IsSubClassOf(c)
	} else {
		return other.IsImplements(c)
	}
}

// IsSubClassOf 判断S是否是T的子类，实际上也就是判断T是否是S的（直接或间接）超类。
func (c *Class) IsSubClassOf(other *Class) bool {
	for super := c.superClass; super != nil; super = super.superClass {
		if super == other {
			return true
		}
	}

	return false
}

// IsImplements 判断c是否实现了iface接口，就看c或c的（直接或间接）超类是否实现了某个接口iface' ,iface’要么是iface，要么是iface的子接口。
func (c *Class) IsImplements(iface *Class) bool {
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

func (c *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(c)
}
