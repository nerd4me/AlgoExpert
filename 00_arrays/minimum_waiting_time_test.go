package arrays

import "sort"

func MinimumWaitingTime(queries []int) int {
	result := 0
	sort.Ints(queries)
	for i, v := range queries {
		timesLeft := len(queries) - (i + 1)
		result += timesLeft * v
	}
	return result
}
