package util

type PriorityQueue[E any] struct {
	queue []E

	size int

	modCount int

	MaxArraySize int
}
