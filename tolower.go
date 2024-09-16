package piscine

func ToLower(s string) string {
	mySlice := []rune(s)
	for i, c := range mySlice {
		if c >= 65 && c <= 90 {
			mySlice[i] = c + 32
		}
	}
	return string(mySlice)
}
