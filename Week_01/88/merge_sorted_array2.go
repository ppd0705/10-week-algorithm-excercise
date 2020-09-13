package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	l := m + n - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[l] = nums1[m]
			m--
		} else {
			nums1[l] = nums2[n]
			n--
		}
		l--
	}
	if n >= 0 {
		copy(nums1[:n+1], nums2[:n+1])
	}
}
