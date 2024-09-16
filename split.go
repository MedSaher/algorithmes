package piscine

func Split(s, sep string) []string {
	str := ""
	mySlice := []string{}
	for i := 0; i < len(s); i++ {
		separator := false
		if i < len(s)-len(sep)+1 {
			if s[i:i+len(sep)] == sep {
				separator = true
			}
		}
		if separator {
			mySlice = append(mySlice, str)
			str = ""
			i += len(sep) - 1
		} else {
			str += string(s[i])
		}
	}
	mySlice = append(mySlice, str)
	return mySlice
}
