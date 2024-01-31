package problems

func SmallerNumbersThanCurrent(nums []int) []int {
	count := 0
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
		ans[i] = count
		count = 0
	}
	return ans
}
