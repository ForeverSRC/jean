package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlag      uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlag:      reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (mi *MemberInfo) AccessFlags() uint16 {
	return mi.accessFlag
}

func (mi *MemberInfo) Name() string {
	return mi.cp.getUtf8(mi.nameIndex)
}

func (mi *MemberInfo) Descriptor() string {
	return mi.cp.getUtf8(mi.descriptorIndex)
}

func (mi *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrType := attrInfo.(type) {
		case *CodeAttribute:
			return attrType
		}
	}

	return nil
}

func (mi *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrType := attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrType
		}
	}

	return nil
}

func (mi *MemberInfo) ExceptionsAttribute() *ExceptionsAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrType := attrInfo.(type) {
		case *ExceptionsAttribute:
			return attrType
		}
	}
	return nil
}

func (mi *MemberInfo) RuntimeVisibleAnnotationsAttributeData() []byte {
	return mi.getUnparsedAttributeData("RuntimeVisibleAnnotations")
}
func (mi *MemberInfo) RuntimeVisibleParameterAnnotationsAttributeData() []byte {
	return mi.getUnparsedAttributeData("RuntimeVisibleParameterAnnotationsAttribute")
}
func (mi *MemberInfo) AnnotationDefaultAttributeData() []byte {
	return mi.getUnparsedAttributeData("AnnotationDefault")
}

func (mi *MemberInfo) getUnparsedAttributeData(name string) []byte {
	for _, attrInfo := range mi.attributes {
		switch unparsedAttr := attrInfo.(type) {
		case *UnparsedAttribute:
			if unparsedAttr.name == name {
				return unparsedAttr.info
			}
		}
	}
	return nil
}
