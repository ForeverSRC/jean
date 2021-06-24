package heap

import "unicode/utf16"

// go string -> Java String
var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C"), chars}

	jStr := loader.LoadClass("java/lang/String").NewObject()
	// hack java.lang.String constructor method
	jStr.SetRefVar("value", "[C", jChars)

	internedStrings[goStr] = jStr
	return jStr

}

func stringToUtf16(s string) []uint16 {
	runes := []rune(s) //utf32
	return utf16.Encode(runes)

}
