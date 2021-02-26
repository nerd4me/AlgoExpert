package arrays

import "testing"

func GetNthFib(n int) int {
	if n <= 1 {
		return 0
	}
	x, y := 0, 1
	for i := 2; i < n-1; i++ {
		x, y = y, x+y
	}
	return y
}

func TestGetNthFib(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(GetNthFib(i))
	}
}
