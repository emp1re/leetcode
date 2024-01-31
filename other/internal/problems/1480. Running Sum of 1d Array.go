package problems

func RunningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j < len(nums)-1; j++ {
			nums[i] += nums[j]
			break
		}
	}
	return nums
}
func runningSum(nums []int) []int {
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] += nums[i] + prefix[i]
	}
	return prefix[1:]
}
