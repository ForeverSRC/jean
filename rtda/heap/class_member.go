package heap

import "jean/classfile"

type ClassMember struct {
	accessFlags    uint16
	name           string
	descriptor     string
	signature      string
	annotationData []byte // RuntimeVisibleAnnotations_attribute
	class          *Class
}

func (cm *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	cm.accessFlags = memberInfo.AccessFlags()
	cm.name = memberInfo.Name()
	cm.descriptor = memberInfo.Descriptor()
}

func (cm *ClassMember) IsPublic() bool {
	return cm.accessFlags&ACC_PUBLIC != 0
}

func (cm *ClassMember) IsPrivate() bool {
	return cm.accessFlags&ACC_PRIVATE != 0
}

func (cm *ClassMember) IsProtected() bool {
	return cm.accessFlags&ACC_PROTECTED != 0
}

func (cm *ClassMember) IsStatic() bool {
	return cm.accessFlags&ACC_STATIC != 0
}

func (cm *ClassMember) IsFinal() bool {
	return cm.accessFlags&ACC_FINAL != 0
}

func (cm *ClassMember) IsSynthetic() bool {
	return cm.accessFlags&ACC_SYNTHETIC != 0
}

// cm能否被other访问
func (cm *ClassMember) isAccessibleTo(d *Class) bool {
	if cm.IsPublic() {
		return true
	}

	c := cm.class
	if cm.IsProtected() {
		return d == c || d.IsSubClassOf(c) || c.GetPackageName() == d.GetPackageName()
	}

	if !cm.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}

	return d == c
}

func (cm *ClassMember) Class() *Class {
	return cm.class
}

func (cm *ClassMember) Name() string {
	return cm.name
}

func (cm *ClassMember) Descriptor() string {
	return cm.descriptor
}

func (cm *ClassMember) AccessFlags() uint16 {
	return cm.accessFlags
}

func (cm *ClassMember) Signature() string {
	return cm.signature
}

func (cm *ClassMember) AnnotationData() []byte {
	return cm.annotationData
}
