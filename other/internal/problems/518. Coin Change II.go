package problems

func Сhange(amount int, coins []int) int {

	set := make([]int, amount+1)
	set[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			set[i] = set[i] + set[i-coin]
		}
	}

	return set[amount]

}
