package arrays

import "testing"

func ValidateSubsequence(array []int, sequence []int) bool {
	arrIdx, seqIdx := 0, 0
	for arrIdx < len(array) && seqIdx < len(sequence) {
		if sequence[seqIdx] == array[arrIdx] {
			seqIdx++
		}
		arrIdx++
	}
	return seqIdx == len(sequence)
}

func TestValidateSubsequence(t *testing.T) {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{5, 1, 22, -1, 10}
	t.Log(ValidateSubsequence(array, sequence))
}
