package problems

func Shuffle(nums []int, n int) []int {
	ans := []int{}

	for i := 0; i < len(nums)/2; i++ {
		ans = append(ans, nums[i])
		ans = append(ans, nums[n])
		n++
	}
	return ans
}
