package searching

func BinarySearch(array []int, target int) int {
	l, r := 0, len(array) - 1
	for l <= r {
		m := (l + r) / 2
		if array[m] == target {
			return m
		} else if array[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
