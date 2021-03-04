package heaps

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap([]int{9, 8, 2, 10, 293, 29, 29, 29, 6})
	length := heap.length()
	for i := 0; i < length; i++ {
		heap.Remove()
		t.Log(*heap)
	}
}
