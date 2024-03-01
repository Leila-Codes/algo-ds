package algo_ds

type Node[T interface{}] struct {
	Value T
	Next  *Node[T]
}

type BinaryNode[T interface{}] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}
