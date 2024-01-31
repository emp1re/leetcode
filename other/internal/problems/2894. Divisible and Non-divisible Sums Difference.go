package problems

func DifferenceOfSums(n int, m int) int {
	sumN := 0
	sumM := 0
	i := 1
	for i <= n {
		if i%m != 0 {
			sumN += i
		}
		if i%m == 0 {
			sumM += i
		}
		i++
	}
	return sumN - sumM
}
