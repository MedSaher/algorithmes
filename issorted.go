package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascend := true
	descend := true
	for i := 1; i < len(a); i++ {
		if f(a[i], a[i-1]) > 0 {
			ascend = false
		} else if f(a[i], a[i-1]) < 0 {
			descend = false
		}
	}
	return ascend || descend
}
