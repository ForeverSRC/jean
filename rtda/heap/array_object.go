package heap

func (o *Object) Bytes() []int8 {
	return o.data.([]int8)
}

func (o *Object) Shorts() []int16 {
	return o.data.([]int16)
}

func (o *Object) Ints() []int32 {
	return o.data.([]int32)
}

func (o *Object) Longs() []int64 {
	return o.data.([]int64)
}

func (o *Object) Chars() []uint16 {
	return o.data.([]uint16)
}

func (o *Object) Floats() []float32 {
	return o.data.([]float32)
}

func (o *Object) Doubles() []float64 {
	return o.data.([]float64)
}

func (o *Object) Refs() []*Object {
	return o.data.([]*Object)
}

func (o *Object) ArrayLength() int32 {
	switch o.data.(type) {
	case []int8:
		return int32(len(o.data.([]int8)))
	case []int16:
		return int32(len(o.data.([]int16)))
	case []int32:
		return int32(len(o.data.([]int32)))
	case []int64:
		return int32(len(o.data.([]int64)))
	case []uint16:
		return int32(len(o.data.([]uint16)))
	case []float32:
		return int32(len(o.data.([]float32)))
	case []float64:
		return int32(len(o.data.([]float64)))
	case []*Object:
		return int32(len(o.data.([]*Object)))
	default:
		panic("Not array!")
	}
}

func ArrayCopy(src *Object, srcPos int32, dest *Object, destPos int32, length int32) {
	switch src.data.(type) {
	case []float32:
		// float
		_src := src.data.([]float32)[srcPos : srcPos+length]
		_dst := dest.data.([]float32)[destPos : destPos+length]
		copy(_dst, _src)
	case []float64:
		// double
		_src := src.data.([]float64)[srcPos : srcPos+length]
		_dst := dest.data.([]float64)[destPos : destPos+length]
		copy(_dst, _src)
	case []int64:
		// long
		_src := src.data.([]int64)[srcPos : srcPos+length]
		_dst := dest.data.([]int64)[destPos : destPos+length]
		copy(_dst, _src)
	case []int32:
		// int
		_src := src.data.([]int32)[srcPos : srcPos+length]
		_dst := dest.data.([]int32)[destPos : destPos+length]
		copy(_dst, _src)
	case []int16:
		// short
		_src := src.data.([]int16)[srcPos : srcPos+length]
		_dst := dest.data.([]int16)[destPos : destPos+length]
		copy(_dst, _src)
	case []int8:
		// byte boolean
		_src := src.data.([]int8)[srcPos : srcPos+length]
		_dst := dest.data.([]int8)[destPos : destPos+length]
		copy(_dst, _src)
	case []uint16:
		// char
		_src := src.data.([]uint16)[srcPos : srcPos+length]
		_dst := dest.data.([]uint16)[destPos : destPos+length]
		copy(_dst, _src)
	case []*Object:
		// ref
		_src := src.data.([]*Object)[srcPos : srcPos+length]
		_dst := dest.data.([]*Object)[destPos : destPos+length]
		copy(_dst, _src)
	default:
		panic("Not array!")
	}
}
