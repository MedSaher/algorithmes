package piscine

import "github.com/01-edu/z01"

const N = 8

var position = [N]int{}

func isSafe(queenNumber, rowPosition int) bool {
	for i := 0; i < queenNumber; i += 2 {
		otherRowPos := position[i]
		if otherRowPos == rowPosition ||
			otherRowPos == rowPosition-(queenNumber-i) ||
			otherRowPos == rowPosition+(queenNumber-i) {
			return false
		}
	}
	return true
}

func solve(k int) {
	if k == N {
		for i := 0; i < N; i++ {
			z01.PrintRune(rune(rune(position[i] + '1')))
		}
		z01.PrintRune(rune('\n'))
	} else {
		for i := 0; i < N; i++ {
			if isSafe(k, i) {
				position[k] = i
				solve(k + 1)
			}
		}
	}
}

func EightQueens() {
	solve(0)
}
