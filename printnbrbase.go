package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	strSlice := []rune(base)
	if len(strSlice) < 2 || LookUp(base, "-") || LookUp(base, "+") || !AreAllCharsUnique(strSlice) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr == -9223372036854775808 { // Edge case for minimum int value
		z01.PrintRune('-')
		PrintNbrBaseHelper(922337203685477580, base)
		z01.PrintRune('8') // Manually handle the last digit
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	PrintNbrBaseHelper(nbr, base)
}

func PrintNbrBaseHelper(nbr int, base string) {
	strSlice := []rune(base)
	baseLen := len(strSlice)
	var myResult []int

	for nbr > 0 {
		myResult = append(myResult, nbr%baseLen)
		nbr /= baseLen
	}

	if len(myResult) == 0 { // If the number is 0
		z01.PrintRune(strSlice[0])
		return
	}

	for i := len(myResult) - 1; i >= 0; i-- {
		z01.PrintRune(strSlice[myResult[i]])
	}
}

func AreAllCharsUnique(slice []rune) bool {
	charMap := make(map[rune]bool)
	for _, char := range slice {
		if _, exists := charMap[char]; exists {
			return false
		}
		charMap[char] = true
	}
	return true
}
