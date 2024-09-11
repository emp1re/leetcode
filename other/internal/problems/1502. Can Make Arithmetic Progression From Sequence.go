package problems

import "slices"

func CanMakeArithmeticProgression(arr []int) bool {
	slices.Sort(arr)

	cache := 0

	for i := 0; i < len(arr)-1; i++ {
		if i == 0 {
			cache = arr[i+1] - arr[i]
		}
		if arr[i+1]-arr[i] != cache {
			return false
		}
	}

	return true
}
