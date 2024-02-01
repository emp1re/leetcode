package problems

func BalancedStringSplit(s string) int {
	var counter int
	var countRL int
	for _, v := range s {
		if string(v) == "L" {
			countRL -= 1
		}
		if string(v) == "R" {
			countRL += 1
		}
		if countRL == 0 {
			counter++
		}
	}
	return counter
}
