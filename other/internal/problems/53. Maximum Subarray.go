package problems

func MaxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := nums[0]
	if len(nums) == 1 {
		return nums[0]
	}
	for i := 1; i < len(nums); i++ {
		curSum += nums[i]

		if curSum < nums[i] {
			curSum = nums[i]
		}
		if curSum > maxSum {
			maxSum = curSum
		}
		if maxSum < curSum {
			maxSum = curSum
		}
	}
	return maxSum
}
