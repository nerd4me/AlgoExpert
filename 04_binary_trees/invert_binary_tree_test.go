package binary_trees

func (tree *BinaryTree) InvertBinaryTree() {
	tree.Left, tree.Right = tree.Right, tree.Left
	if tree.Left != nil {
		tree.Left.InvertBinaryTree()
	}
	if tree.Right != nil {
		tree.Right.InvertBinaryTree()
	}
}
