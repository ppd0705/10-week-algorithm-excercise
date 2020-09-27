package main

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func getKthMagicNumber2(k int) int {
	intHeap := &IntHeap{1}
	heap.Init(intHeap)
	visited := map[int]bool{1: true}
	factors := [3]int{3, 5, 7}
	n := 1
	i := 0
	for i < k {
		i++
		n = (heap.Pop(intHeap)).(int)
		for j := 0; j < len(factors); j++ {
			m := n * factors[j]
			if !visited[m] {
				heap.Push(intHeap, m)
				visited[m] = true
			}
		}
	}
	return n
}
