package problems

import "slices"

func MaxWidthOfVerticalArea(points [][]int) int {
	x := []int{}
	for k, _ := range points {
		x = append(x, points[k][0])
	}
	slices.Sort(x)
	max := 0
	for i := 0; i < len(x)-1; i++ {
		width := x[i+1] - x[i]
		if max < width {
			max = width
		}
	}
	return max
}
