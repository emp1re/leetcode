package problems

func NumIdenticalPairs(nums []int) int {
	var count int
	for k, v := range nums {
		for i := 1; i < len(nums); i++ {
			if v == nums[i] && k < i {
				count++
			} else {
				continue
			}
		}
	}
	return count
}
