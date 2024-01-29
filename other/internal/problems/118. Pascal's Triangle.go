package problems

func Generate(numRows int) [][]int {
	var ans [][]int
	for i := 1; i <= numRows; i++ {
		set := make([]int, i)
		if len(ans) == 0 {
			set[0] = 1
		} else {
			for k := 0; k <= len(ans[i-2]); k++ {
				if k == 0 {
					set[k] = 1
				} else if k == i-1 {
					set[k] = 1
				} else {
					set[k] = ans[i-2][k-1] + ans[i-2][k]

				}
			}

		}
		ans = append(ans, set)
	}
	return ans
}
