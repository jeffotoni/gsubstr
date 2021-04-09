package gsubstr

// substr (text, 0, 4)
// substr (text, 2, 2)
func Substr(value string, leni, lenf int) string {
	if leni < 0 || lenf < 0 {
		return ""
	}
	if len(value) < leni {
		return ""
	}
	leny := lenf
	soma := leni + lenf
	if soma > len(value) {
		soma = len(value)
	}
	leny = soma
	return value[leni:leny]
}
