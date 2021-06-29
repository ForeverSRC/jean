package lang

import (
	"jean/constants"
	"jean/instructions/base"
	"jean/native"
	"jean/rtda/heap"
	"jean/rtda/jvmstack"
	"os"
	"runtime"
	"time"
)

//public static native void arraycopy(Object src,  int  srcPos, Object dest, int destPos, int length);
// Javadoc:
//Throws:
//IndexOutOfBoundsException – if copying would cause access of data outside array bounds.
//ArrayStoreException – if an element in the src array could not be stored into the dest array because of a type mismatch.
//NullPointerException – if either src or dest is null.
func arrayCopy(frame *jvmstack.Frame) {
	vars := frame.LocalVars()

	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)

	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}

	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}

	if srcPos < 0 || destPos < 0 || length < 0 || srcPos+length > src.ArrayLength() || destPos+length > dest.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	heap.ArrayCopy(src, srcPos, dest, destPos, length)
}

func checkArrayCopy(src, dest *heap.Object) bool {
	srcClass := src.Class()
	destClass := dest.Class()

	if !srcClass.IsArray() || !destClass.IsArray() {
		return false
	}

	if srcClass.ComponentClass().IsPrimitive() || destClass.ComponentClass().IsPrimitive() {
		return srcClass == destClass
	}

	return true
}

//private static native Properties initProperties(Properties prop);
func initProperties(frame *jvmstack.Frame) {
	vars := frame.LocalVars()
	props := vars.GetRef(0)

	stack := frame.OperandStack()
	stack.PushRef(props)

	// public synchronized Object setProperty(String key,String value)
	setPropertyMethod := props.Class().GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	thread := frame.Thread()
	for key, val := range sysProps {
		jKey := heap.JString(frame.Method().Class().Loader(), key)
		jVal := heap.JString(frame.Method().Class().Loader(), val)
		ops := jvmstack.NewOperandStack(3)
		ops.PushRef(props)
		ops.PushRef(jKey)
		ops.PushRef(jVal)
		shimFrame := jvmstack.NewShimFrame(thread, ops)
		thread.PushFrame(shimFrame)
		base.InvokeMethod(shimFrame, setPropertyMethod)

	}
}

// private static native void setIn0(InputStream in);
// (Ljava/io/InputStream;)V
func setIn0(frame *jvmstack.Frame) {
	vars := frame.LocalVars()
	in := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetRefVar("in", "Ljava/io/InputStream;", in)
}

// private static native void setOut0(PrintStream out);
func setOut0(frame *jvmstack.Frame) {
	out := frame.LocalVars().GetRef(0)
	sysClass := frame.Method().Class()
	sysClass.SetRefVar("out", "Ljava/io/PrintStream;", out)
}

// private static native void setErr0(PrintStream err);
// (Ljava/io/PrintStream;)V
func setErr0(frame *jvmstack.Frame) {
	vars := frame.LocalVars()
	err := vars.GetRef(0)

	sysClass := frame.Method().Class()
	sysClass.SetRefVar("err", "Ljava/io/PrintStream;", err)
}

// public static native long currentTimeMillis();
// ()J
func currentTimeMillis(frame *jvmstack.Frame) {
	millis := time.Now().UnixNano() / int64(time.Millisecond)
	stack := frame.OperandStack()
	stack.PushLong(millis)
}

func init() {
	native.Register(constants.JavaLangSystem,
		"arraycopy",
		"(Ljava/lang/Object;ILjava/lang/Object;II)V",
		arrayCopy)
	native.Register(constants.JavaLangSystem,
		"initProperties",
		"(Ljava/util/Properties;)Ljava/util/Properties;",
		initProperties)
	native.Register(constants.JavaLangSystem, "setIn0", "(Ljava/io/InputStream;)V", setIn0)
	native.Register(constants.JavaLangSystem, "setOut0", "(Ljava/io/PrintStream;)V", setOut0)
	native.Register(constants.JavaLangSystem, "setErr0", "(Ljava/io/PrintStream;)V", setErr0)
	native.Register(constants.JavaLangSystem, "currentTimeMillis", "()J", currentTimeMillis)
}

var sysProps = map[string]string{
	"java.version":         "1.8.0",
	"java.vendor":          "jean",
	"java.vendor.url":      "https://github.com/ForeverSRC/jean",
	"java.home":            "todo",
	"java.class.version":   "52.0",
	"java.class.path":      "todo",
	"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
	"os.name":              runtime.GOOS,   // todo
	"os.arch":              runtime.GOARCH, // todo
	"os.version":           "",             // todo
	"file.separator":       string(os.PathSeparator),
	"path.separator":       string(os.PathListSeparator),
	"line.separator":       "\n", // todo
	"user.name":            "",   // todo
	"user.home":            "",   // todo
	"user.dir":             ".",  // todo
	"user.country":         "CN", // todo
	"file.encoding":        "UTF-8",
	"sun.stdout.encoding":  "UTF-8",
	"sun.stderr.encoding":  "UTF-8",
}
