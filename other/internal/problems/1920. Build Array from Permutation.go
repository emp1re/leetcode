package problems

func BuildArray(nums []int) []int {
	ans := []int{}
	for k, _ := range nums {
		ans = append(ans, nums[nums[k]])
	}
	return ans
}
