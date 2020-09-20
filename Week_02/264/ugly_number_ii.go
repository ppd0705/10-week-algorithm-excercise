package main

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func nthUglyNumber(n int) int {
	intHeap := &IntHeap{1}
	heap.Init(intHeap)

	curr := 1
	seen := map[int]bool{1: true}
	factors := []int{2, 3, 5}
	for i := 0; i < n; i++ {
		curr = heap.Pop(intHeap).(int)
		for _, f := range factors {
			num := f * curr
			if _, ok := seen[num]; !ok {
				heap.Push(intHeap, num)
				seen[num] = true
			}
		}
	}
	return curr
}
