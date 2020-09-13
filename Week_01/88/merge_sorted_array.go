package main

func merge0(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	for j := 0; j < n; j++ {
		for nums1[i] <= nums2[j] && i < m+j {
			i++
		}
		copy(nums1[i+1:m+j+1], nums1[i:m+j])
		nums1[i] = nums2[j]
		i++
	}
}
