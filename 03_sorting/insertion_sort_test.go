package sorting

import "testing"

func InsertionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		j := i
		for j >= 0 && array[j] > array[j+1] {
			swap(array, j, j+1)
			j--
		}
	}
	return array
}

func swap(array []int, i int, j int) {
	array[i], array[j] = array[j], array[i]
}

func TestInsertionSort(t *testing.T) {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	t.Log(InsertionSort(array))
}
