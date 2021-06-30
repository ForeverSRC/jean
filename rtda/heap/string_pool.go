package heap

import (
	"jean/constants"
	"unicode/utf16"
)

// go string -> Java String
var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	chars := stringToUtf16(goStr)
	jChars := &Object{
		class: loader.LoadClass("[C"),
		data:  chars,
	}

	jStr := loader.LoadClass(constants.JavaLangString).NewObject()
	// hack java.lang.String constructor method
	jStr.SetRefVar("value", "[C", jChars)

	internedStrings[goStr] = jStr
	return jStr

}

func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

func stringToUtf16(s string) []uint16 {
	runes := []rune(s) //utf32
	return utf16.Encode(runes)

}

func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s)
	return string(runes)
}

func InternString(jStr *Object) *Object {
	goStr := GoString(jStr)
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	internedStrings[goStr] = jStr
	return jStr
}
