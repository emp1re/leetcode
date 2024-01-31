package problems

func MaximumWealth(accounts [][]int) int {
	var max int
	for k, _ := range accounts {
		var sum int
		for i := 0; i < len(accounts[k]); i++ {
			sum += accounts[k][i]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}
