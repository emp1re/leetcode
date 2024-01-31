package problems

func productExceptSelf(nums []int) []int {
	pref := make([]int, len(nums))
	for k := range pref {
		pref[k] = 1
		for i := 0; i < len(nums); i++ {

			if k == i {
				pref[k] = pref[k] * 1
			} else {
				pref[k] *= nums[i]
			}
		}

	}
	return pref
}
