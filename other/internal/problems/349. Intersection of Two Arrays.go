package problems

import "slices"

func Intersection(nums1 []int, nums2 []int) []int {
	var inter []int
	set := make(map[int]int)
	for _, v := range nums1 {
		set[v] = v
	}
	for _, v := range nums2 {
		if _, ok := set[v]; ok {
			if !slices.Contains(inter, v) {
				inter = append(inter, v)
			}
		}
	}
	return inter
}
