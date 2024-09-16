package piscine

func Capitalize(s string) string {
	mySlice := []rune(s)
	if len(mySlice) == 0 {
		return s
	}

	if mySlice[0] >= 'a' && mySlice[0] <= 'z' {
		mySlice[0] -= 'a' - 'A'
	}

	for i := 1; i < len(mySlice); i++ {
		if isLower(mySlice[i]) {
			if !isLetterOrDigit(mySlice[i-1]) {
				mySlice[i] -= 'a' - 'A'
			}
		} else if isUpper(mySlice[i]) {
			if isLetterOrDigit(mySlice[i-1]) {
				mySlice[i] += 'a' - 'A'
			}
		}
	}
	return string(mySlice)
}

func isLetterOrDigit(c rune) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func isLower(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func isUpper(c rune) bool {
	return c >= 'A' && c <= 'Z'
}
