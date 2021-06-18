package heap

type SymbolicRef struct {
	rtCp      *ConstantPool
	className string

	// cache
	class *Class
}

func (s *SymbolicRef) ResolvedClass() *Class {
	if s.class == nil {
		s.resolveClassRef()
	}
	return s.class
}

func (s *SymbolicRef) resolveClassRef() {
	d := s.rtCp.class
	c := d.loader.LoadClass(s.className)
	if !c.IsAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	s.class = c
}
