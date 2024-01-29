package problems

func GroupThePeople(groupSizes []int) [][]int {
	var grs [][]int
	hash := map[int][]int{}
	for k, v := range groupSizes {
		hash[v] = append(hash[v], k)

	}
	for k, v := range hash {
		for i := 0; i < len(v); i += k {
			grs = append(grs, v[i:i+k])
		}
	}
	return grs
}
