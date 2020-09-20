package main

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return []int{}
	}
	heap := arr[:k:k]
	for i := k/2 - 1; i >= 0; i-- {
		heapify(heap, i, k-1)
	}

	for i := k; i < len(arr); i++ {
		if arr[i] < heap[0] {
			heap[0] = arr[i]
			heapify(heap, 0, k-1)
		}
	}
	return heap
}

func heapify(heap []int, root, end int) {
	for {
		child := root*2 + 1 // left child
		if child > end {
			return
		}
		if child+1 <= end && heap[child] < heap[child+1] {
			child++ // right child
		}
		if heap[root] > heap[child] {
			return
		}
		heap[root], heap[child] = heap[child], heap[root]
		root = child
	}
}
