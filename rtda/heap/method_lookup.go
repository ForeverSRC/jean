package heap

func LookupMethodInClass(class *Class, name, descriptor string) *Method {
	for c := class; c != nil; c = class.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}

	return nil
}

func lookupMethodInInterface(ifaces []*Class, name, descriptor string) *Method {
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := lookupMethodInInterface(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}

	return nil

}
