package parser

type GoTypes struct {
	types map[string]string // map[syntax]style
}

func NewGoTypes() *GoTypes {
	return &GoTypes{
		types: make(map[string]string),
	}
}

func (gt *GoTypes) AddType(t, s string) {
	gt.types[t] = s
}

func (gt *GoTypes) GetType(t string) string {
	return gt.types[t]
}

func GetGoTypes() *GoTypes {
	goTypes := NewGoTypes()
	goTypes.AddType("import", "text-red")
	goTypes.AddType("const", "text-red")
	goTypes.AddType("package", "text-red")
	goTypes.AddType("func", "text-red")
	goTypes.AddType("defer", "text-red")
	goTypes.AddType("return", "text-red")
	goTypes.AddType("if", "text-red")
	goTypes.AddType("else", "text-red")
	goTypes.AddType("for", "text-red")
	goTypes.AddType("range", "text-red")
	goTypes.AddType("switch", "text-red")
	goTypes.AddType("case", "text-red")
	goTypes.AddType("default", "text-red")
	goTypes.AddType("break", "text-red")
	goTypes.AddType("continue", "text-red")
	goTypes.AddType(":=", "text-red")
	goTypes.AddType("=", "text-red")
	goTypes.AddType("==", "text-red")
	goTypes.AddType("!=", "text-red")
	goTypes.AddType(">", "text-red")
	goTypes.AddType("<", "text-red")
	goTypes.AddType(">=", "text-red")
	goTypes.AddType("<=", "text-red")
	goTypes.AddType("&&", "text-red")
	goTypes.AddType("||", "text-red")
	goTypes.AddType("!", "text-red")
	goTypes.AddType("+", "text-red")
	goTypes.AddType("-", "text-red")
	goTypes.AddType("*", "text-red")
	goTypes.AddType("/", "text-red")
	goTypes.AddType("%", "text-red")
	goTypes.AddType("++", "text-red")
	goTypes.AddType("--", "text-red")
	goTypes.AddType("<<", "text-red")
	goTypes.AddType(">>", "text-red")
	goTypes.AddType("&", "text-red")
	goTypes.AddType("|", "text-red")
	goTypes.AddType("^", "text-red")
	goTypes.AddType("&^", "text-red")
	goTypes.AddType("+=", "text-red")

	goTypes.AddType("string", "text-blue")
	goTypes.AddType("int", "text-blue")
	goTypes.AddType("int8", "text-blue")
	goTypes.AddType("int16", "text-blue")
	goTypes.AddType("int32", "text-blue")
	goTypes.AddType("int64", "text-blue")
	goTypes.AddType("uint", "text-blue")
	goTypes.AddType("uint8", "text-blue")
	goTypes.AddType("uint16", "text-blue")
	goTypes.AddType("uint32", "text-blue")
	goTypes.AddType("uint64", "text-blue")
	goTypes.AddType("uintptr", "text-blue")
	goTypes.AddType("float32", "text-blue")
	goTypes.AddType("float64", "text-blue")
	goTypes.AddType("complex64", "text-blue")
	goTypes.AddType("complex128", "text-blue")
	goTypes.AddType("byte", "text-blue")
	goTypes.AddType("rune", "text-blue")
	goTypes.AddType("bool", "text-blue")
	goTypes.AddType("error", "text-blue")
	goTypes.AddType("nil", "text-blue")
	goTypes.AddType("true", "text-blue")
	goTypes.AddType("false", "text-blue")
	goTypes.AddType("iota", "text-blue")
	goTypes.AddType("new", "text-blue")
	goTypes.AddType("make", "text-blue")

	goTypes.AddType("String()", "text-green")
	goTypes.AddType("Println", "text-green")
	goTypes.AddType("Printf", "text-green")
	goTypes.AddType("Print", "text-green")
	goTypes.AddType("Sprint", "text-green")
	goTypes.AddType("Sprintf", "text-green")
	goTypes.AddType("Sprintln", "text-green")
	goTypes.AddType("Fprint", "text-green")
	goTypes.AddType("Fprintf", "text-green")
	goTypes.AddType("Fprintln", "text-green")
	goTypes.AddType("Errorf", "text-green")
	goTypes.AddType("Fatal", "text-green")
	goTypes.AddType("Fatalf", "text-green")
	goTypes.AddType("Fatalln", "text-green")
	goTypes.AddType("Panic", "text-green")
	goTypes.AddType("Panicf", "text-green")

	goTypes.AddType("fmt", "text-text")
	goTypes.AddType("log", "text-text")
	goTypes.AddType("os", "text-text")
	goTypes.AddType("io", "text-text")
	goTypes.AddType("bufio", "text-text")
	goTypes.AddType("errors", "text-text")
	goTypes.AddType("strings", "text-text")
	goTypes.AddType("strconv", "text-text")
	goTypes.AddType("time", "text-text")
	goTypes.AddType("math", "text-text")
	goTypes.AddType("encoding", "text-text")
	goTypes.AddType("encoding/json", "text-text")
	goTypes.AddType("encoding/xml", "text-text")
	goTypes.AddType("encoding/csv", "text-text")

	return goTypes
}
