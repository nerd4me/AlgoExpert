package arrays

import "testing"

func FirstDuplicateValue(array []int) int {
	for _, v := range array {
		absValue := abs(v)
		if array[absValue-1] < 0 {
			return absValue
		}
		array[absValue-1] *= -1
	}
	return -1
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}

func TestFirstDuplicateValue(t *testing.T) {
	array := []int{2, 1, 5, 2, 3, 3, 4}
	t.Log(FirstDuplicateValue(array))
}
