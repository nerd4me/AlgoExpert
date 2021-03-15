package binary_trees

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NewBinaryTree(root int, values ...int) *BinaryTree {
	tree := &BinaryTree{Value: root}
	tree.Insert(values, 0)
	return tree
}

func (tree *BinaryTree) Insert(values []int, i int) *BinaryTree {
	if i >= len(values) {
		return tree
	}
	val := values[i]
	queue := []*BinaryTree{tree}
	for len(queue) > 0 {
		var current *BinaryTree
		current, queue = queue[0], queue[1:]
		if current.Left == nil {
			current.Left = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Left)

		if current.Right == nil {
			current.Right = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Right)
	}

	tree.Insert(values, i+1)
	return tree
}
