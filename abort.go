package piscine

func Abort(a, b, c, d, e int) int {
	mySlice := []int{a, b, c, d, e}
	mySlice = SortIntegerTable(mySlice)
	return mySlice[2]
}
