package problems

func MaxArea(height []int) int {
	left := 0
	right := len(height) - 1
	volume := 0
	for left < right {
		if height[left] < height[right] {
			v := (right - left) * height[left]
			if v > volume {
				volume = v
			}
			left++
		} else if height[left] > height[right] {
			v := (right - left) * height[right]
			if v > volume {
				volume = v
			}
			right--
		} else {
			v := (right - left) * height[right]
			if v > volume {
				volume = v
			}
			left++
		}
	}
	return volume
}
