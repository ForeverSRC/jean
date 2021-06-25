package heap

func (c *Class) NewArray(count uint) *Object {
	if !c.IsArray() {
		panic("Not array class: " + c.name)
	}

	switch c.name {
	case "[Z":
		return &Object{class: c, data: make([]int8, count)}
	case "[B":
		return &Object{class: c, data: make([]int8, count)}
	case "[C":
		return &Object{class: c, data: make([]uint16, count)}
	case "[S":
		return &Object{class: c, data: make([]int16, count)}
	case "[I":
		return &Object{class: c, data: make([]int32, count)}
	case "[J":
		return &Object{class: c, data: make([]int64, count)}
	case "[F":
		return &Object{class: c, data: make([]float32, count)}
	case "[D":
		return &Object{class: c, data: make([]float64, count)}
	default:
		return &Object{class: c, data: make([]*Object, count)}
	}
}

func (c *Class) IsArray() bool {
	return c.name[0] == '['
}

func (c *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(c.name)
	return c.loader.LoadClass(componentClassName)
}
