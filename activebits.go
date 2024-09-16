package piscine

func ActiveBits(n int) int {
	bnbr := ConvertBase(itoa(n), "0123456789", "01")
	counter := 0
	for _, c := range bnbr {
		if c == '1' {
			counter++
		}
	}
	return counter
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var digits []byte
	for n > 0 {
		digit := n % 10
		digits = append(digits, byte('0'+digit))
		n /= 10
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}
