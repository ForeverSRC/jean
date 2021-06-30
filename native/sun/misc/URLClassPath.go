package misc

import (
	"jean/native"
	"jean/rtda/jvmstack"
)

func init() {
	native.Register("sun/misc/URLClassPath", "getLookupCacheURLs", "(Ljava/lang/ClassLoader;)[Ljava/net/URL;", getLookupCacheURLs)
}

// private static native URL[] getLookupCacheURLs(ClassLoader var0);
// (Ljava/lang/ClassLoader;)[Ljava/net/URL;
func getLookupCacheURLs(frame *jvmstack.Frame) {
	frame.OperandStack().PushRef(nil)
}
