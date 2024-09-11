package problems

func MoveZeroes(nums []int) []int {
	nulls := 0
	i := 0
	for _, v := range nums {
		if v == 0 {
			nulls += 1
			continue
		}
		nums[i] = v
		i++
	}

	if nulls != 0 {
		for j := len(nums) - nulls; j < len(nums); j++ {
			nums[j] = 0
		}
	}
	return nums

}
