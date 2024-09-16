package piscine

func Unmatch(a []int) int {
	myMap := make(map[int]int)
	for _, num := range a {
		myMap[num]++
	}

	for _, num := range a {
		if myMap[num]%2 != 0 {
			return num
		}
	}
	return -1
}
