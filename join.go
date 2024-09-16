package piscine

func Join(strs []string, sep string) string {
	var str string
	for i, char := range strs {
		str += char
		if i != len(strs)-1 {
			str += sep
		}
	}
	return str
}
