package problems

func GetConcatenation(nums []int) []int {

	for _, v := range nums {
		nums = append(nums, v)
	}
	return nums
}
