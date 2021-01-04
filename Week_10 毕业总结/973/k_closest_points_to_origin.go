package main

func kClosest(points [][]int, K int) [][]int {
	n := len(points)
	if K == n {
		return points
	}
	heap := points[:K]
	for i := K /2 -1; i >= 0;i-- {
		heapify(heap, i, K-1)
	}

	for i := K; i < n;i++ {
		if isCloser(points[i], heap[0]) {
			heap[0] = points[i]
			heapify(heap, 0, K-1)
		}
	}
	return heap
}


func isCloser(x , y []int) bool {
	if x[0] * x[0] + x[1] * x[1] <= y[0] * y[0] + y[1] * y[1] {
		return true
	}
	return false
}


func heapify(points [][]int, root, end int) {
	for {
		child := root*2+1
		if child > end {
			return
		}
		if child < end && isCloser(points[child], points[child+1]) {
			child++
		}

		if isCloser(points[child], points[root]) {
			return
		}
		points[root], points[child] = points[child], points[root]
		root = child
	}
}