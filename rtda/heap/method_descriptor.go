package heap

type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

func (md *MethodDescriptor) addParameterType(t string) {
	pLen := len(md.parameterTypes)
	if pLen == cap(md.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, md.parameterTypes)
		md.parameterTypes = s
	}

	md.parameterTypes = append(md.parameterTypes, t)
}
