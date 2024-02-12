package problems

func CountDigits(num int) int {
	var count int
	if num < 10 {
		return 1
	}
	//copy *num* for search digits!!!!
	cp := num
	x := 1
	for x != 0 {
		// x take last number
		x = cp % 10
		// divide number on 10, for take next digits
		cp /= 10
		if x == 0 {
			return count
		}

		if num%x == 0 {
			count++
		}
	}
	return count
}
