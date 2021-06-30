package heap

type Object struct {
	class *Class
	data  interface{}
	extra interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  NewSlots(class.instanceSlotCount),
	}
}

func (o *Object) Fields() Slots {
	return o.data.(Slots)
}

func (o *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(o.class)
}

func (o *Object) Class() *Class {
	return o.class
}

func (o *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := o.class.getField(name, descriptor, false)
	slots := o.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (o *Object) GetRefVar(name, descriptor string) *Object {
	field := o.class.getField(name, descriptor, false)
	slots := o.data.(Slots)
	return slots.GetRef(field.slotId)
}

func (o *Object) SetIntVar(name, descriptor string, val int32) {
	field := o.class.getField(name, descriptor, false)
	slots := o.data.(Slots)
	slots.SetInt(field.slotId, val)
}
func (o *Object) GetIntVar(name, descriptor string) int32 {
	field := o.class.getField(name, descriptor, false)
	slots := o.data.(Slots)
	return slots.GetInt(field.slotId)
}

func (o *Object) Data() interface{} {
	return o.data
}

func (o *Object) SetExtra(extra interface{}) {
	o.extra = extra
}

func (o *Object) Extra() interface{} {
	return o.extra
}
