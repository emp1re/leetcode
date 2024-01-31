package problems

func FindWordsContaining(words []string, x byte) []int {
	ans := []int{}
	for k, v := range words {
		for _, value := range v {
			if byte(value) == x {
				ans = append(ans, k)
				break
			} else {
				continue
			}
		}
	}
	return ans
}
