package problems

func DecompressRLElist(nums []int) []int {
	set := []int{}
	for i := 0; i < len(nums); i = i + 2 {
		for k := nums[i]; k > 0; k-- {
			set = append(set, nums[i+1])
		}
	}
	return set
}
