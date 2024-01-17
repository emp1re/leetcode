package problems

func TwoSum(numbers []int, target int) []int {
	n := len(numbers)
	left, right := 0, n-1
	answer := []int{}
	for left < right {
		if (numbers[left] + numbers[right]) > target {
			right--
		}
		if (numbers[left] + numbers[right]) < target {
			left++
		}
		if (numbers[left] + numbers[right]) == target {
			return []int{left + 1, right + 1}
		}
	}
	return answer
}
