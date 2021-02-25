package arrays

import "sort"

func ClassPhotos(redShirtHeights[]int, blueShirtHeights[] int) bool {
	sort.Sort(sort.Reverse(sort.IntSlice(redShirtHeights)))
	sort.Sort(sort.Reverse(sort.IntSlice(blueShirtHeights)))

	redShirtFirst := 0
	if redShirtHeights[0] < blueShirtHeights[0] {
		redShirtFirst = 1
	}
	for i:=0; i < len(redShirtHeights);i++ {
		redHeight, blueHeight := redShirtHeights[i], blueShirtHeights[i]
		if redShirtFirst == 1 {
			if redHeight >= blueHeight {
				return false
			}
		} else {
			if blueHeight >= redHeight {
				return false
			}
		}
	}
	return true
}
