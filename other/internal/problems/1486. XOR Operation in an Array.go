package problems

func XorOperation(n int, start int) int {
	nums := []int{}
	var xor int
	for i := 0; i < n; i++ {
		nums = append(nums, start+2*i)
		xor ^= nums[i]
	}

	return xor
}
