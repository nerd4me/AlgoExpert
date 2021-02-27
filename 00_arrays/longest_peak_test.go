package arrays

import (
	"testing"
)

func LongestPeak(array []int) int {
	longestPeakLength := 0
	i := 1
	for i < len(array)-1 {
		isPeak := array[i-1] < array[i] && array[i] > array[i+1]
		if !isPeak {
			i += 1
			continue
		}
		leftIdx := i - 2
		for leftIdx >= 0 && array[leftIdx] < array[leftIdx+1] {
			leftIdx--
		}
		rightIdx := i + 2
		for rightIdx < len(array) && array[rightIdx] < array[rightIdx-1] {
			rightIdx++
		}
		currentPeakLength := rightIdx - leftIdx - 1
		if currentPeakLength > longestPeakLength {
			longestPeakLength = currentPeakLength
		}
		i = rightIdx
	}
	return longestPeakLength
}

func TestLongestPeak(t *testing.T) {
	arr := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
	t.Log(LongestPeak(arr))
}
