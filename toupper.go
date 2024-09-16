package piscine

func ToUpper(s string) string {
	mySlice := []rune(s)
	for i, c := range mySlice {
		if c >= 97 && c <= 122 {
			mySlice[i] = c - 32
		}
	}
	return string(mySlice)
}
