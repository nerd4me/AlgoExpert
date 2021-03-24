package dynamic_programming

import "testing"

func NumberOfWaysToMakeChange(n int, denoms []int) int {
	ways := make([]int, n+1)
	ways[0] = 1
	for _, d := range denoms {
		for i := 1; i < len(ways); i++ {
			if d <= i {
				ways[i] += ways[i-d]
			}
		}
	}
	return ways[n]
}

func TestNumberOfWaysToMakeChange(t *testing.T) {
	change := NumberOfWaysToMakeChange(25, []int{1, 5, 10, 20})
	t.Log(change)
}
