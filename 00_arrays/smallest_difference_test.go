package arrays

import (
	"sort"
	"testing"
)

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	i1, i2 := 0, 0
	result := make([]int, 2, 2)
	sDiff := 1<<63 - 1
	for i1 < len(array1) && i2 < len(array2) {
		v1, v2 := array1[i1], array2[i2]
		if v1 == v2 {
			return []int{v1, v2}
		} else {
			var diff int
			if v1 < v2 {
				diff = v2 - v1
				i1++
			} else {
				diff = v1 - v2
				i2++
			}
			if diff < sDiff {
				sDiff = diff
				result[0] = v1
				result[1] = v2
			}
		}
	}
	return result
}

func TestSmallestDifference(t *testing.T) {
	arr1 := []int{-1, 5, 10, 20, 28, 3}
	arr2 := []int{26, 134, 135, 15, 17}
	r := SmallestDifference(arr1, arr2)
	t.Log(r)
}
