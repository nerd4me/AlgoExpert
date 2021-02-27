package arrays

import "testing"

func SpiralTraverse(array [][]int) []int {
	startRow, endRow := 0, len(array)-1
	startCol, endCol := 0, len(array[0])-1
	result := make([]int, 0, (endRow+1)*(endCol+1))
	for startRow <= endRow && startCol <= endCol {
		for c := startCol; c <= endCol; c++ {
			result = append(result, array[startRow][c])
		}
		for c := startRow + 1; c <= endRow; c++ {
			result = append(result, array[c][endCol])
		}
		for c := endCol - 1; c >= startCol; c-- {
			if startRow == endRow {
				break
			}
			result = append(result, array[endRow][c])
		}
		for c := endRow - 1; c > startRow; c-- {
			if startCol == endCol {
				break
			}
			result = append(result, array[c][startCol])
		}
		startRow++
		endRow--
		startCol++
		endCol--
	}
	return result
}

func TestSpiralTraverse(t *testing.T) {
	array := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}
	t.Log(SpiralTraverse(array))
}
