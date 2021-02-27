package arrays

import "testing"

// O(n) time | O(n) space
func ArrayOfProducts(array []int) []int {
	products := make([]int, len(array))

	leftRunningProduct := 1
	for i := range array {
		products[i] = leftRunningProduct
		leftRunningProduct *= array[i]
	}

	rightRunningProduct := 1
	for i := len(array) - 1; i >= 0; i-- {
		products[i] *= rightRunningProduct
		rightRunningProduct *= array[i]
	}
	return products
}

func TestArrayOfProducts(t *testing.T) {
	array := []int{5, 1, 4, 2}
	t.Log(ArrayOfProducts(array))
}
