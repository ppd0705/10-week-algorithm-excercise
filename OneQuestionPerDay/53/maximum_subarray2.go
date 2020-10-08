package main

type status struct {
	lSum int
	rSum int
	iSum int
	mSum int
}

func pushUp(l, r status) status {
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, l.rSum+r.iSum)
	iSum := l.iSum + r.iSum
	mSum := max(max(l.mSum, r.mSum), l.rSum + r.lSum)
	return status{lSum, rSum, iSum, mSum}
}

func get(arr []int, l, r int) status {
	if l == r {
		return status{arr[l], arr[l], arr[l], arr[l]}
	}
	m := (l + r) / 2
	left := get(arr, l, m)
	right := get(arr, m+1, r)
	return pushUp(left, right)
}
func maxSubArray2(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}
