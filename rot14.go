package piscine

func Rot14(s string) string {
	x := []rune(s)
	y := make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		if x[i] >= 'a' && x[i] <= 'z' {
			y[i] = ((x[i] - 'a' + rune(14)) % 26) + 'a'
		} else if x[i] >= 'A' && x[i] <= 'Z' {
			y[i] = ((x[i] - 'A' + rune(14)) % 26) + 'A'
		} else {
			y[i] = x[i]
		}
	}
	return string(y)
}
