package binary_trees

import "testing"


func BranchSums(root *BinaryTree) []int {
	var sums []int
	calculateBranchSums(root, 0, &sums)
	return sums
}

func calculateBranchSums(node *BinaryTree, runningSum int, sums *[]int) {
	if node == nil {
		return
	}

	runningSum += node.Value
	if node.Left == nil && node.Right == nil {
		*sums = append(*sums, runningSum)
		return
	}

	calculateBranchSums(node.Left, runningSum, sums)
	calculateBranchSums(node.Right, runningSum, sums)
}

func TestBranchSums(t *testing.T) {
	tree := NewBinaryTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	//expected := []int{15, 16, 18, 10, 11}
	output := BranchSums(tree)
	t.Log(output)
}
