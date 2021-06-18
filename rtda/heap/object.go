package heap

type Object struct {
	class  *Class
	fields Slots
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: NewSlots(class.instanceSlotCount),
	}
}

func (o *Object) Fields() Slots {
	return o.fields
}

func (o *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(o.class)
}
