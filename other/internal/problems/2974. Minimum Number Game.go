package problems

import "slices"

func NumberGame(nums []int) []int {
	slices.Sort(nums)
	i := 0
	for i < len(nums)-1 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
		i += 2
	}
	return nums
}
