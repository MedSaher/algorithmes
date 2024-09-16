package piscine

func SplitWhiteSpaces(s string) []string {
	mySlice := []string{}
	var str string
	for _, c := range s {
		if c == ' ' || c == '\n' || c == '\t' {
			if str != "" {
				mySlice = append(mySlice, str)
				str = ""
			}
		} else {
			str += string(c)
		}
	}
	if str != "" {
		mySlice = append(mySlice, str)
	}
	return mySlice
}
