package sb

var ostr string

func CreateNew() {
	ostr = ""
}

func Appen(nstr string) {
	ostr += nstr
}

func Appenl(nstr string) {
	ostr += "\n" + nstr
}

func AppenEmptyl() {
	ostr += "\n"
}

func Retrieve() string {
	return ostr
}
