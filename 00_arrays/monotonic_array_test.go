package arrays

import "testing"

const (
	ASC = iota // 升序
	DSC        // 降序
)

func IsMonotonic(array []int) bool {
	if len(array) <= 2 {
		return true
	}
	current := array[0]
	order := DSC
	if current <= array[len(array)-1] {
		order = ASC
	}
	for i := 1; i < len(array); i++ {
		if order == ASC && current > array[i] {
			return false
		} else if order == DSC && current < array[i] {
			return false
		}
		current = array[i]
	}
	return true
}

func TestIsMonotonic(t *testing.T) {
	array := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}
	t.Log(IsMonotonic(array))
}
