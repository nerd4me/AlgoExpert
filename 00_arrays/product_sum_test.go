package arrays

import "testing"

type SpecialArray []interface{}

func ProductSum(array SpecialArray) int {
	return helper(array, 1)
}

func helper(array SpecialArray, multiplier int) int {
	sum := 0
	for _, el := range array {
		if cast, ok := el.(SpecialArray); ok {
			sum += helper(cast, multiplier+1)
		} else if cast, ok := el.(int); ok {
			sum += cast
		}
	}
	return multiplier * sum
}

func TestProductSum(t *testing.T) {
	array := SpecialArray{
		5, 2,
		SpecialArray{7, -1},
		3,
		SpecialArray{
			6,
			SpecialArray{
				-13, 8,
			},
			4,
		},
	}
	sum := ProductSum(array)
	t.Log(sum)
}