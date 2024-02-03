package problems

func NumberOfSteps(num int) int {
	var count int
	for num != 0 {
		if num%2 == 0 {
			count++
			num = num >> 1
			if num == 0 {
				return count
			}
		} else {
			num = num - 1
			count++
		}
	}
	return count
}
