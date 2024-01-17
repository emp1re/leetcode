package internal

func PrefixSum(arr []int) []int {
	prefixSum := make([]int, len(arr))
	prefixSum[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
		
	}
	return prefixSum
}
func PrefixSumInterval(arr []int, start, end int) int {
	if start == 0 {
		return arr[end]
	}
	return arr[end] - arr[start-1]
}
