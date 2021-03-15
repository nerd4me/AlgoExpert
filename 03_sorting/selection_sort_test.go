package sorting

import "testing"

func SelectionSort(array []int) []int {
	for i, v := range array {
		min := v
		for j := i+1; j < len(array); j++ {
			if min > array[j] {
				min = array[j]
			}
		}
	}
	return nil
}

func TestSelectionSort(t *testing.T) {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	t.Log(SelectionSort(array))
}
