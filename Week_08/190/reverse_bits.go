package main

func reverseBits(num uint32) uint32 {
	pow, ans := uint32(31), uint32(0)
	for num > 0 {
		ans += (1 & num) << pow
		pow--
		num >>= 1
	}
	return ans
}