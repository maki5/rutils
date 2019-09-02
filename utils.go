package rutils

func IntPtr(i int) *int {
	return &i
}

func InverseInt(n int) int {
	return n - (n * 2)
}

// just a copy of strings Blank created to avoid import cycle on internal usage 
// TODO: refactor it later
func Blank(str string) bool {
	if len(str) == 0 {
		return true
	}

	for _, s := range str {
		if string(s) != " " {
			return false
		}
	}

	return true
}
