package arrays

import (
	"sort"
	"testing"
)

func NonConsturctibleChange(coins []int) int {
	sort.Ints(coins)
	change := 0
	for _, v := range coins {
		if v > change + 1 {
			return change + 1
		}
		change += v
	}
	return change + 1
}

func TestNonConstructibleChange(t *testing.T) {
	coins := []int{1, 3, 5, 23, 11, 1}
	t.Log(NonConsturctibleChange(coins))
}
