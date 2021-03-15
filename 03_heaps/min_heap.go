package heaps

type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap()
	return ptr
}

func (h *MinHeap) BuildHeap() {
	first := (h.length() - 2) / 2
	for currentIdx := first + 1; currentIdx >= 0; currentIdx-- {
		h.siftDown(currentIdx, h.length() - 1)
	}
}

func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	h.siftUp()
}

func (h *MinHeap) Remove() int {
	l := h.length()
	h.swap(0, l-1)
	peeked := (*h)[l-1]
	*h = (*h)[0 : l-1]
	h.siftDown(0, l-2)
	return peeked
}

// O(log(n)) time | O(1) space
func (h *MinHeap) siftUp() {
	currentIdx := h.length() - 1
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 {
		current, parent := (*h)[currentIdx], (*h)[parentIdx]
		if current < parent {
			h.swap(currentIdx, parentIdx)
			currentIdx = parentIdx
			parentIdx = (currentIdx - 1) / 2
		} else {
			return
		}
	}
}

func (h *MinHeap) siftDown(currentIdx, endIdx int) {
	leftIdx := currentIdx*2 + 1
	for leftIdx <= endIdx {
		rightIdx := -1
		if currentIdx*2+2 <= endIdx {
			rightIdx = currentIdx*2 + 2
		}
		idxToSwap := leftIdx
		if rightIdx > -1 && (*h)[rightIdx] < (*h)[leftIdx] {
			idxToSwap = rightIdx
		}
		if (*h)[idxToSwap] < (*h)[currentIdx] {
			h.swap(idxToSwap, currentIdx)
			currentIdx = idxToSwap
			leftIdx = currentIdx*2 + 1
		} else {
			return
		}
	}
}

func (h MinHeap) swap(i1, i2 int) {
	h[i1], h[i2] = h[i2], h[i1]
}

func (h MinHeap) length() int {
	return len(h)
}

