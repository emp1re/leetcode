package problems

func NumJewelsInStones(jewels string, stones string) int {
	hmap := make(map[rune]int)
	var count int
	for _, v := range jewels {
		hmap[v] = 0
	}
	for _, v := range stones {
		_, ok := hmap[v]
		if !ok {
			continue
		} else {
			count++
		}

	}
	return count
}
