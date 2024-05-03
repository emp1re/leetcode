package problems

// https://leetcode.com/problems/number-of-arithmetic-triplets/description/
func ArithmeticTriplets(nums []int, diff int) int {
	var c int
	for i := 1; i < len(nums); i++ {
		for k := 0; k < i; k++ {
			if nums[i]-nums[k] == diff {
				for j := i; j <= len(nums)-1; j++ {
					if nums[j]-nums[i] == diff {
						c++
					}
				}
			}
		}
	}
	return c
}
