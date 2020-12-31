package main

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	s := sum(nums)
	if s%2 == 1 {
		return false
	}
	target := s / 2

	// dp[i][j] represent whether can select x numbers from nums[0:i+1] and their sum equals j
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, target+1)
	}
	for i := 0; i < len(nums); i++ {
		dp[i][0] = true
	}
	if nums[0] <= target {
		dp[0][nums[0]] = true
	}

	for i := 1; i < len(nums); i++ {
		for j := 1; j <= target; j++ {
			if nums[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[len(nums)-1][target]
}

func sum(nums []int) int {
	s := 0
	for _, i := range nums {
		s += i
	}
	return s
}
