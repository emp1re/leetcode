package problems

import "math"

func ContainsNearbyDuplicate(nums []int, k int) bool {
	window := 1

	for window < k+1 {
		for i := 0; i < len(nums)-window; i++ {
			j := i + window
			if nums[i] == nums[j] && math.Abs(float64(i-j)) <= float64(k) {
				return true
			}
		}
		window++
	}

	return false
}
