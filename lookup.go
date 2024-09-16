package piscine

func LookUp(s string, toFind string) bool {
	lenS := len(s)
	lenToFind := len(toFind)

	if lenToFind == 0 {
		return false
	}
	if lenToFind > lenS {
		return false
	}
	for i := 0; i <= lenS-lenToFind; i++ {
		if s[i:i+lenToFind] == toFind {
			return true
		}
	}
	return false
}
