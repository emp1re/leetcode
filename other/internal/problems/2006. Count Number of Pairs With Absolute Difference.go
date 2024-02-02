package problems

func CountKDifference(nums []int, k int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if abs(nums[i], nums[j]) == k {
				count++
			} else {
				continue
			}
		}
	}
	return count
}
func abs(a, b int) int {
	if a-b < 0 {
		return -(a - b)
	} else {
		return a - b
	}

}
