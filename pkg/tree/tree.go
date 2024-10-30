package tree

// A generic N-ary tree. All Function are recursive which might lead to stack overflows.
type Tree[T any] struct {
	Value    T
	parent   *Tree[T]
	children []*Tree[T]
}

// Get all of the children of the tree node. Order is the Append order.
func Children[T any](tree *Tree[T]) []*Tree[T] {
	return tree.children
}

// Get the parrent of the tree node.
func Parent[T any](tree *Tree[T]) *Tree[T] {
	return tree.parent
}

// Add new child to the tree node.
func Append[T any](tree *Tree[T], val T) {
	tree.children = append(tree.children, &Tree[T]{Value: val, parent: tree})
}

// Collect all nodes in the order of post-order DFS.
func PostorderCollect[T any](tree *Tree[T], out *[]T) {
	PostorderCollectFilter(tree, func(t T) bool { return true }, out)
}

// Collect all nodes in the order of post-order DFS filtering any nodes for which the filter
// function returns false.
func PostorderCollectFilter[T any](tree *Tree[T], filter func(T) bool, out *[]T) {
	for _, v := range tree.children {
		PostorderCollectFilter(v, filter, out)
	}
	if filter(tree.Value) {
		*out = append(*out, tree.Value)
	}
}

// Call the clb function on every node in the order of post-order DFS.
func PostorderCallback[T any](tree *Tree[T], clb func(*Tree[T])) {
	for _, v := range tree.children {
		PostorderCallback(v, clb)
	}
	clb(tree)
}
