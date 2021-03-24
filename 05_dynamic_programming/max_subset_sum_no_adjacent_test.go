package dynamic_programming

import "testing"

// O(n) time | O(1) space
func MaxSubsetSumNoAdjacent(array []int) int {
	if len(array) == 0 {
		return 0
	}
	if len(array) == 1 {
		return array[0]
	}
	first := max(array[0], array[1])
	second := array[0]
	for i := 2; i < len(array); i++ {
		current := max(first, second+array[i])
		second, first = first, current
	}
	return first
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestMaxSubsetSumNoAdjacent(t *testing.T) {
	array := []int{75, 105, 120, 75, 90, 135}
	t.Log(MaxSubsetSumNoAdjacent(array))
}
