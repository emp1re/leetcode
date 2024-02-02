package problems

func Trap(height []int) int {
	maxW, inx := 0, 0
	nowM := 0
	water := 0
	for k, val := range height {
		if maxW < val {
			maxW = val
			inx = k
		}
	}
	for i := 0; i <= inx; i++ {
		if height[i] > nowM {
			nowM = height[i]
		}
		water += nowM - height[i]
	}
	nowM = 0
	for i := len(height) - 1; i > inx; i-- {
		if height[i] > nowM {
			nowM = height[i]
		}
		water += nowM - height[i]
	}
	return water
}
