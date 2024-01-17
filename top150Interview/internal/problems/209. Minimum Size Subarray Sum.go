package problems

func MinSubArrayLen(target int, nums []int) int {
	window := 1
	var sum int
	for window <= len(nums) {
		for i := 0; i < window; i++ {
			sum += nums[i]
			if sum >= target {
				return window
			}
		}
		for i := window; i < len(nums); i++ {
			sum += nums[i] - nums[i-window]
			if sum >= target {
				return window
			}
		}
		window++
		sum = 0
	}

	return 0

}
