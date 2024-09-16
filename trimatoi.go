package piscine

func TrimAtoi(s string) int {
	digits := []int{}
	var result int
	var isNegative bool
	for _, c := range s {
		if c == '-' {
			if len(digits) == 0 {
				isNegative = true
			}
		}
		if IsNumeric(string(c)) {
			digits = append(digits, int(c))
		}
	}
	for _, n := range digits {
		result = result*10 + int(n-'0')
	}
	if isNegative {
		return -result
	}
	return result
}
