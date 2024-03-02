package graph

type Node[T comparable] struct {
	Value       T
	Connections Edge[T]
}

type Edge[T comparable] struct {
	Destination *Node[T]
	//Weight comparable
}
