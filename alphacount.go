package piscine

func AlphaCount(s string) int {
	var counter int
	for _, char := range s {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			counter++
		}
	}
	return counter
}
