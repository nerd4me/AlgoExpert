package arrays

import (
	"sort"
	"testing"
)

func ThreeNumbersSum(nums []int, targetSum int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var result [][]int
	for i, v := range nums {
		s := targetSum - v
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum < s {
				left++
			} else if sum > s {
				right--
			} else {
				result = append(result, []int{v, nums[left], nums[right]})
				left++
				right--
			}
		}
	}
	return result
}

func TestThreeNumbersSum(t *testing.T) {
	nums := []int{12, 3, 1, 2, -6, 5, -8, 6}
	t.Log(ThreeNumbersSum(nums, 0))
}
