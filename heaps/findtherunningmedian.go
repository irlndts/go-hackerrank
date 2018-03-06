package heaps

import "container/heap"

// Find The Running Median
// https://www.hackerrank.com/challenges/find-the-running-median/problem
// Heaps

// AddDelay adds a delay to the slice
func (m *Median) AddDelay(delay int) {
	heap.Push(m.MaxHeap, delay)
	heap.Push(m.MinHeap, m.MaxHeap.Top())
	heap.Pop(m.MaxHeap)

	if m.MaxHeap.Len() < m.MinHeap.Len() {
		heap.Push(m.MaxHeap, m.MinHeap.Top())
		heap.Pop(m.MinHeap)
	}
}

// GetMedian returns a median value of the slice
func (m Median) GetMedian() float64 {
	if m.MinHeap.Len() < m.MaxHeap.Len() {
		return float64(m.MaxHeap.Top())
	}

	return float64(m.MaxHeap.Top()+m.MinHeap.Top()) * 0.5
}

type Median struct {
	MaxHeap *MaxHeap
	MinHeap *MinHeap
	Size    int
}

func InitMedianQueue(size int) *Median {
	return &Median{
		Size:    size,
		MaxHeap: NewMaxHeap(),
		MinHeap: NewMinHeap(),
	}
}

// MEDIAN interfaces realizations
// MaxHeap is a max heap of ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Top() int {
	a := *h
	if len(a) != 0 {
		return int(a[0])
	}
	return 0
}

func NewMaxHeap() *MaxHeap {
	pq := &MaxHeap{}
	heap.Init(pq)
	return pq
}

// MinHeap is a min heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinHeap) Top() int {
	a := *h
	if len(a) != 0 {
		return int(a[0])
	}
	return 0
}

func NewMinHeap() *MinHeap {
	pq := &MinHeap{}
	heap.Init(pq)
	return pq
}
