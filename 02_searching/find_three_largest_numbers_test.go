package searching

import (
	"math"
	"testing"
)

func FindThreeLargestNumbers(array []int) []int {
	three := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	for _, v := range array {
		if v > three[2] {
			shiftArray(three, 2, v)
		} else if v > three[1] {
			shiftArray(three, 1, v)
		} else if v > three[0] {
			shiftArray(three, 0, v)
		}
	}
	return three
}

func shiftArray(toShift []int, index, num int) {
	for i := 0; i < index; i++ {
		toShift[i] = toShift[i+1]
	}
	toShift[index] = num
}

func TestFindThreeLargestNumbers(t *testing.T) {
	array := []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	t.Log(FindThreeLargestNumbers(array))
}
