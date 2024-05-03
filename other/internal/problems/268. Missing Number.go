package problems

import "sort"

func MissingNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	i := 0
	for _, v := range nums {
		if v != i {

			return i
		}
		i++
	}
	return nums[len(nums)-1] + 1
}
