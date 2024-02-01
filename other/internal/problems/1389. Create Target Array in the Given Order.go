package problems

func CreateTargetArray(nums []int, index []int) []int {
	arr := []int{}
	for k, v := range nums {
		if len(arr) > index[k] {
			simulation := make([]int, len(arr))
			copy(simulation, arr)
			left := simulation[:index[k]]
			right := arr[index[k]:]
			left = append(left, 1)
			left[index[k]] = v
			arr = append(left, right...)
		} else {
			arr = append(arr, v)
		}
	}

	return arr
}
