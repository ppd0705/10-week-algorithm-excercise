package main

func countBits(num int) []int {
	res := make([]int, num+1)
	for i:=1;i <= num; i++ {
		//res[i] = i % 2 + res[i /2]
		res[i] = i & 1 + res[i >> 1]
	}
	return res
}