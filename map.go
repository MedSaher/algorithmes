package piscine

func Map(f func(int) bool, a []int) []bool {
	boolMap := make(map[int]bool)
	for _, c := range a {
		boolMap[c] = f(c)
	}
	boolSlice := []bool{}
	for _, d := range a {
		boolSlice = append(boolSlice, boolMap[d])
	}
	return boolSlice
}
